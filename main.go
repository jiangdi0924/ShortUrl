package main

import (
	"fmt"
	"short/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type ShortUrl struct {
	ID uint `gorm:"primarykey"`

	Url       string
	Short     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func main() {
	app := fiber.New()
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		fmt.Println("[DATABASE]::CONNECTION_ERROR")
		panic(err)
	}

	db.AutoMigrate(&ShortUrl{})

	app.Get("/set/*", func(c *fiber.Ctx) error {
		long_url := c.Params("*")
		long := &ShortUrl{Url: long_url}

		result := db.Create(&long)

		if result.Error != nil {
			return fiber.NewError(fiber.StatusConflict, result.Error.Error())
		}

		short := fmt.Sprintf("%v/t%v", utils.Env("DOMAIN", "https://url.cscript.site"), long.ID)

		db.Model(&ShortUrl{}).
			Where(&ShortUrl{ID: long.ID}).
			Update("short", short)

		record := &ShortUrl{}
		db.Model(&ShortUrl{}).
			Where("url = ?", long_url).
			First(record)

		return c.JSON(fiber.Map{
			"CScriptUrl": short,
		})
	})

	app.Get("/t:id", func(c *fiber.Ctx) error {
		record := &ShortUrl{}
		db.Model(&ShortUrl{}).
			Where("id", c.Params("id")).
			First(record)
		// fmt.Println(record.Url)
		c.Redirect(record.Url)

		return c.JSON(record)
	})

	app.Get("/lists", func(c *fiber.Ctx) error {
		records := []ShortUrl{}

		db.Model(&ShortUrl{}).
			Find(&records)
		return c.JSON(records)
	})

	app.Listen(fmt.Sprintf(":%v", utils.Env("PORT", "3000")))
}

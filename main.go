package main

import (
	"fmt"
	"strconv"
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
		db.Model(&ShortUrl{}).
			Where(&ShortUrl{ID: long.ID}).
			Update("short", c.Hostname()+strconv.Itoa(int(long.ID)))
		return c.JSON(long)
	})

	app.Get("/goto/*", func(c *fiber.Ctx) error {
		record := &ShortUrl{}
		db.Model(&ShortUrl{}).
			Where("short", c.Params("*")).
			First(record)
		fmt.Println(record.Url)
		c.Redirect(record.Url)
		return nil
	})

	app.Listen(fmt.Sprintf(":%v", 3003))
}

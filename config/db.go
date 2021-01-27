package config

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBConnect() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		fmt.Println("[DATABASE]::CONNECTION_ERROR")
		panic(err)
	}

	// Setting the database connection to use in routes
	DB = db

	fmt.Println("[DATABASE]::CONNECTED")
}
func Migrate(tables ...interface{}) error {
	return DB.AutoMigrate(tables...)
}

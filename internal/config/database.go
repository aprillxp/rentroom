package config

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("./internal/database/rentroom.db?_foreign_keys=on"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database", err)
	}

	DB = database
}

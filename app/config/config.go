package config

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	database *gorm.DB
)

func Connect() {
	connection, err := gorm.Open(sqlite.Open("database/database.sqlite"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	database = connection
}

func GetDatabase() *gorm.DB {
	return database
}

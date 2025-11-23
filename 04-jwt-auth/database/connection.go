package database

import (
	"jwt-auth/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	connection, err := gorm.Open(sqlite.Open("user.db"), &gorm.Config{})
	if err != nil {
		panic("Could not connect to the database")
	}
	DB = connection

	// migrations
	connection.AutoMigrate(models.User{})
}

// Package models/setup provides an interface to connect to the database

package models

import (
	"github.com/jinzhu/gorm"
)

// DB variable represents our connection to the database
var DB *gorm.DB

// ConnectDatabase opens a sqllite3 database connection, and migrates the created models
func ConnectDatabase() {
	database, err := gorm.Open("sqlite3", "loans.db")

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&Loan{})

	DB = database
}

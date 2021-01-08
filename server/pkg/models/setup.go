// Package models/setup provides an interface to connect to the database

package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// DB variable represents our connection to the database
var DB *gorm.DB

// ConnectDatabase opens a sqllite3 database connection, and migrates the created models
func ConnectDatabase() {
	db, err := gorm.Open(sqlite.Open("loans.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(&Loan{})

	DB = db
}

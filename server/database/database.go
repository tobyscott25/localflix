package database

import (
	"fmt"
	"log"

	"localflix/server/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Database *gorm.DB

func Connect(libraryLocation string) {

	// Use SQLite database, create if not exists
	db, dbErr := gorm.Open(sqlite.Open(libraryLocation+"/library.db"), &gorm.Config{})
	if dbErr != nil {
		log.Fatal("❌ Failed to find, connect to and/or create a SQLite database")
	}
	Database = db

	fmt.Println("🚀 Successfully connected to the database and ran any outstanding migrations")
}

func RunMigrations() {

	err := Database.AutoMigrate(&model.Video{})

	if err != nil {
		log.Fatal("❌ Failed to run migrations")
	}
	fmt.Println("🚀 Successfully ran all database migrations")
}

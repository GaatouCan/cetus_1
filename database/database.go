package database

import (
	model2 "demo/model"
	"gorm.io/gorm"
	"log"
)

func InitDatabaseTables(db *gorm.DB) {
	if err := db.AutoMigrate(&model2.User{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	if err := db.AutoMigrate(&model2.GroceryItem{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
}

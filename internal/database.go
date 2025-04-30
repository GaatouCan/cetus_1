package internal

import (
	"demo/internal/model"
	"gorm.io/gorm"
	"log"
)

func InitDatabaseTables(db *gorm.DB) {
	if err := db.AutoMigrate(&model.User{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	if err := db.AutoMigrate(&model.GroceryItem{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
}

package model

import "gorm.io/gorm"

type GroceryItem struct {
	gorm.Model
	ID       string `json:"id"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
	Category string `json:"category"`
}

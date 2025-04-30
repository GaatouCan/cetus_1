package handler

import (
	"demo/internal/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var groceryItems = []model.GroceryItem{
	{ID: "OP3keqGThTCqmkQoBbx", Name: "Banana", Quantity: 1, Category: "fruit"},
}

type GroceryItemHandler struct {
}

func (h *GroceryItemHandler) GetGroceryItems(c *gin.Context) {
	c.JSON(200, groceryItems)
}

func (h *GroceryItemHandler) CreateGroceryItem(c *gin.Context) {
	var groceryItem model.GroceryItem
	if err := c.ShouldBindJSON(&groceryItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(groceryItem)
	groceryItems = append(groceryItems, groceryItem)

	c.JSON(http.StatusCreated, groceryItem)
}

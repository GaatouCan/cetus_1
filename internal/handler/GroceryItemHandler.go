package handler

import (
	"demo/internal/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var groceryItems = []model.GroceryItem{
	{ID: "OP3keqGThTCqmkQoBbx", Name: "Banana", Quantity: 1, Category: "Fruit"},
	{ID: "OP3keqGfwesdgq", Name: "Milk", Quantity: 1, Category: "Dairy"},
	{ID: "dfaewgwedsew", Name: "Beef Steak", Quantity: 5, Category: "Meat"},
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

func (h *GroceryItemHandler) UpdateGroceryItem(c *gin.Context) {
	id := c.Param("id")

	var groceryItem model.GroceryItem
	if err := c.ShouldBindJSON(&groceryItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	index := -1
	for i, item := range groceryItems {
		if item.ID == id {
			index = i
			break
		}
	}

	if index == -1 {
		groceryItems = append(groceryItems, groceryItem)
		c.JSON(http.StatusCreated, groceryItem)
	} else {
		groceryItems[index] = groceryItem
		c.JSON(http.StatusOK, groceryItem)
	}
}

func (h *GroceryItemHandler) DeleteGroceryItem(c *gin.Context) {
	id := c.Param("id")

	ret := false
	for index, item := range groceryItems {
		if item.ID == id {
			groceryItems = append(groceryItems[:index], groceryItems[index+1:]...)
			ret = true
			break
		}
	}

	if ret {
		c.JSON(http.StatusOK, gin.H{"message": "delete item success"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "item not exist"})
	}
}

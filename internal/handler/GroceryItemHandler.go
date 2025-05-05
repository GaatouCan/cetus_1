package handler

import (
	"demo/internal/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type GroceryItemHandler struct {
	DB *gorm.DB
}

func (h *GroceryItemHandler) GetGroceryItems(c *gin.Context) {
	var result []model.GroceryItem
	if err := h.DB.Find(&result).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *GroceryItemHandler) CreateGroceryItem(c *gin.Context) {
	var groceryItem model.GroceryItem
	if err := c.ShouldBindJSON(&groceryItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.DB.Create(&groceryItem).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, groceryItem)
}

func (h *GroceryItemHandler) UpdateGroceryItem(c *gin.Context) {
	id := c.Param("id")

	var oldGroceryItem model.GroceryItem

	var groceryItem model.GroceryItem
	if err := c.ShouldBindJSON(&groceryItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if groceryItem.ID != id {
		c.JSON(http.StatusBadRequest, gin.H{"message": "id<UNK>"})
		return
	}

	if err := h.DB.Where("id = ?", id).First(&oldGroceryItem).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	if err := h.DB.Updates(&groceryItem).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, groceryItem)
}

func (h *GroceryItemHandler) DeleteGroceryItem(c *gin.Context) {
	id := c.Param("id")
	var groceryItem model.GroceryItem

	if err := h.DB.Where("id = ?", id).First(&groceryItem).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Database not found"})
		return
	}

	if err := h.DB.Delete(&groceryItem).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

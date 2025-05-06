package handler

import (
	"demo/internal/model"
	"errors"
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
	var item model.GroceryItem
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.DB.Create(&item).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, item)
}

func (h *GroceryItemHandler) UpdateGroceryItem(c *gin.Context) {
	var item model.GroceryItem

	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var oldItem model.GroceryItem
	err := h.DB.Where("id = ?", item.ID).First(&oldItem).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = h.DB.Create(&item).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	} else {
		err = h.DB.Updates(&item).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, item)
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

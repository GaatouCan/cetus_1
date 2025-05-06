package handler

import (
	"demo/internal/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type UserHandler struct {
	DB *gorm.DB
}

func (h *UserHandler) GetAllUsers(c *gin.Context) {
	// page := c.Param("page")

	var users []model.User

	if err := h.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	var result []model.UserResponse
	for _, user := range users {
		result = append(result, model.UserResponse{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
		})
	}

	c.JSON(http.StatusOK, result)
}

func (h *UserHandler) GetUser(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	var user model.User
	if err := h.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

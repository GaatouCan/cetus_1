package internal

import (
	"demo/internal/handler"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRouter(router *gin.Engine, db *gorm.DB) {
	// Hello
	userGroup := router.Group("/user")
	{
		h := handler.HelloHandler{}

		router.GET("/hello", h.HelloWorld)

		userGroup.GET("", h.GetUser)
		userGroup.POST("", h.CreateUser)
		userGroup.PUT("", h.UpdateUser)
		userGroup.DELETE("", h.DeleteUser)
	}

	// GroceryItem
	groceryGroup := router.Group("/groceryItem")

	{
		h := handler.GroceryItemHandler{DB: db}

		groceryGroup.GET("", h.GetGroceryItems)
		groceryGroup.POST("", h.CreateGroceryItem)
		groceryGroup.PUT("", h.UpdateGroceryItem)
		groceryGroup.DELETE("/:id", h.DeleteGroceryItem)
	}
}

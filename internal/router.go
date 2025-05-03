package internal

import (
	handler2 "demo/internal/handler"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRouter(router *gin.Engine, db *gorm.DB) {
	// Hello
	userGroup := router.Group("/user")
	{
		handler := handler2.HelloHandler{}

		router.GET("/hello", handler.HelloWorld)

		userGroup.GET("/", handler.GetUser)
		userGroup.POST("/", handler.CreateUser)
		userGroup.PUT("/", handler.UpdateUser)
		userGroup.DELETE("/", handler.DeleteUser)
	}

	// GroceryItem
	groceryGroup := router.Group("/groceryItem")

	{
		handler := handler2.GroceryItemHandler{DB: db}

		groceryGroup.GET("/", handler.GetGroceryItems)
		groceryGroup.POST("/", handler.CreateGroceryItem)
		groceryGroup.PUT("/", handler.UpdateGroceryItem)
		groceryGroup.DELETE("/", handler.DeleteGroceryItem)
	}
}

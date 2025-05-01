package router

import (
	"demo/handler"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRouter(router *gin.Engine, db *gorm.DB) {
	// Hello
	helloHandler := handler.HelloHandler{}

	router.GET("/hello", helloHandler.HelloWorld)

	userGroup := router.Group("/user")
	{
		userGroup.GET("/", helloHandler.GetUser)
		userGroup.POST("/", helloHandler.CreateUser)
		userGroup.PUT("/", helloHandler.UpdateUser)
		userGroup.DELETE("/", helloHandler.DeleteUser)
	}

	// GroceryItem
	groceryItemHandler := handler.GroceryItemHandler{DB: db}

	groceryGroup := router.Group("/groceryItem")
	{
		groceryGroup.GET("/", groceryItemHandler.GetGroceryItems)
		groceryGroup.POST("/", groceryItemHandler.CreateGroceryItem)
		groceryGroup.PUT("/", groceryItemHandler.UpdateGroceryItem)
		groceryGroup.DELETE("/", groceryItemHandler.DeleteGroceryItem)
	}
}

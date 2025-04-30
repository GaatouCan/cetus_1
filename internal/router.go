package internal

import (
	"demo/internal/handler"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRouter(router *gin.Engine, db *gorm.DB) {
	// Hello
	helloHandler := handler.HelloHandler{}

	router.GET("/hello", helloHandler.HelloWorld)
	router.GET("/user", helloHandler.GetUser)
	router.POST("/user", helloHandler.CreateUser)
	router.PUT("/user", helloHandler.UpdateUser)
	router.DELETE("/user", helloHandler.DeleteUser)

	// GroceryItem

	groceryItemHandler := handler.GroceryItemHandler{DB: db}

	router.GET("/groceryItem", groceryItemHandler.GetGroceryItems)
	router.POST("/groceryItem", groceryItemHandler.CreateGroceryItem)
	router.PUT("/groceryItem", groceryItemHandler.UpdateGroceryItem)
	router.DELETE("/groceryItem", groceryItemHandler.DeleteGroceryItem)
}

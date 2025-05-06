package internal

import (
	"demo/internal/auth"
	"demo/internal/handler"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRouter(router *gin.Engine, db *gorm.DB) {
	{
		h := handler.HelloHandler{}

		router.GET("/hello", h.HelloWorld)
	}

	// User
	userGroup := router.Group("/user")
	{
		h := handler.UserHandler{DB: db}

		router.POST("/login", h.UserLogin)
		router.POST("/register", h.CreateUser)

		userGroup.Use(auth.TokenMiddleware())

		userGroup.GET("/page/:page", h.GetAllUsers)
		userGroup.GET("", h.GetUserInfo)
		userGroup.PUT("", h.UpdateUser)
		userGroup.DELETE("/:id", h.DeleteUser)
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

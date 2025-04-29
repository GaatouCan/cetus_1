package main

import (
	"demo/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	helloHandler := handler.HelloHandler{}

	router.GET("/hello", helloHandler.HelloWorld)
	router.GET("/user", helloHandler.GetUser)
	router.POST("/user", helloHandler.CreateUser)

	err := router.Run(":8080")
	if err != nil {
		return
	}
}

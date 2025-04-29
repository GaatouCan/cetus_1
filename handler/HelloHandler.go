package handler

import (
	"demo/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HelloHandler struct{}

func (h *HelloHandler) HelloWorld(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		name = "World"
	}
	c.JSON(http.StatusOK, gin.H{"Greeting": "Hello, " + name})
}

func (h *HelloHandler) GetUser(c *gin.Context) {
	c.JSON(http.StatusOK, model.User{
		ID:       1001,
		Username: "Gaatou",
		Email:    "jiataochen22@gmail.com"})
}

func (h *HelloHandler) CreateUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	fmt.Println(user)
	c.JSON(http.StatusOK, user)
}

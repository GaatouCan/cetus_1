package handler

import (
	"demo/internal/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
		return
	}
	fmt.Println(user)
	c.JSON(http.StatusOK, user)
}

func (h *HelloHandler) UpdateUser(c *gin.Context) {
	idStr := c.Query("id")
	fmt.Println(idStr)

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: Find Target User From Database

	var newUser model.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if newUser.ID != id {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is different"})
		return
	}

	// TODO: Update New User To Database

	c.JSON(http.StatusOK, newUser)
}

func (h *HelloHandler) DeleteUser(c *gin.Context) {
	idStr := c.Query("id")
	fmt.Println(idStr)

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: Delete Target User From Database
	message := fmt.Sprintf("Deleted user[%d]", id)
	c.JSON(http.StatusOK, gin.H{"message": message})
}

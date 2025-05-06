package handler

import (
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

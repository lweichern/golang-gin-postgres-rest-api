package middleware

import (
	"example/http-server/controller"

	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	// Get auth token from header
	token := c.GetHeader("Authorization")

	// Check if token is valid
	if _, ok := controller.Users[token]; ok {
		c.Next()
		return
	}

	c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
	c.Abort()
}
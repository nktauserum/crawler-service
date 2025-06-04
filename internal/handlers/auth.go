package handlers

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func CheckAPIToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		expectedToken := os.Getenv("API_TOKEN")

		// Если API_TOKEN не установлен, считаем что аутентификация в целом отключена
		if expectedToken == "" {
			c.Next()
			return
		}

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization header required",
			})
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid authorization format. Use 'Bearer <token>'",
			})
			c.Abort()
			return
		}

		token := parts[1]
		if token != expectedToken {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid API token",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

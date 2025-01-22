package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func JWT(authSvc *Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is required"})
			c.Abort()
			return
		}

		// Remove Bearer prefix if present
		if strings.HasPrefix(token, "Bearer ") {
			token = token[7:]
		}

		claims, err := authSvc.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Set("user_id", claims.UserID)
		c.Set("role", claims.Role)
		c.Next()
	}
}

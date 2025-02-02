package middleware

import (
	"api-integration/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		userId, err := utils.ValidateToken(tokenString)
		if err != nil || userId == 0 {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}

		c.Set("user_id", userId)
		c.Next()
	}
}

package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/rujool11/chirp-core-service/internal/utils"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.GetHeader("x-jwt-token")

		if tokenStr == "" {
			c.AbortWithStatusJSON(401, gin.H{
				"error": "Missing token",
			})
			return // c.Abort.. stops middleware chain, but function continues running. So we use return here
		}

		userID, err := utils.ValidateJWT(tokenStr)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{
				"error": "Invalid or expired token",
			})
			return
		}

		// store user id in gin context
		c.Set("user_id", userID)
		c.Next()
	}
}

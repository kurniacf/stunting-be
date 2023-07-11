package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kurniacf/stunting-be/pkg/helper"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearerToken, err := helper.ValidateBearerToken(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": err.Error(),
			})
			return
		}

		token, err := helper.GetBearerToken(bearerToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": err.Error(),
			})
			return
		}

		claims, err := helper.VerifyToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			return
		}

		c.Set("user_id", claims["user_id"])

		c.Next()
	}
}

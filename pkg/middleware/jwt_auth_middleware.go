package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearerToken := c.GetHeader("Authorization")

		isTokenValid := strings.HasPrefix(bearerToken, "Bearer ")
		if !isTokenValid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		}

		tokenString := strings.Split(bearerToken, " ")[1]
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		}

		token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if jwt.GetSigningMethod("HS256") != token.Method {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		c.Set("email", claims["email"])

		c.Next()
	}
}

package helper

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func generateToken(claims jwt.Claims, secretKey string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secretKey))

	return tokenString, err
}

func GenerateToken(userId uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	}

	token, err := generateToken(claims, os.Getenv("JWT_SECRET"))

	return token, err
}

func ValidateBearerToken(c *gin.Context) (string, error) {
	bearerToken := c.Request.Header.Get("Authorization")
	if bearerToken == "" {
		return "", fmt.Errorf("Unauthorized")
	}

	isTokenValid := strings.HasPrefix(bearerToken, "Bearer ")
	if !isTokenValid {
		return "", fmt.Errorf("Unauthorized")
	}

	return bearerToken, nil
}

func GetBearerToken(bearerToken string) (string, error) {
	token := strings.Split(bearerToken, " ")[1]
	if token == "" {
		return "", fmt.Errorf("Invalid token")
	}

	return token, nil
}

func VerifyToken(tokenString string) (jwt.MapClaims, error) {
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("Token verification error")
	}

	return claims, nil
}

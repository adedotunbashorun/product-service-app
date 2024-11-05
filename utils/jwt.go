package utils

import (
	"errors"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var jwtSecret = []byte("your_secret_key")

type Claims struct {
	UserID uint   `json:"userID"`
	Role   string `json:"role"`
	jwt.StandardClaims
}

func GenerateJWT(userID uint, roleName string) (string, error) {
	claims := Claims{
		UserID: userID,
		Role:   roleName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// ValidateToken checks the token for validity and returns the claims.
func ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}

func ExtractUserIDFromContext(c *gin.Context) (uint, error) {
	userToken, exists := c.Get("user")
	if !exists {
		return 0, errors.New("user not found in context")
	}

	token, ok := userToken.(*jwt.Token)
	if !ok {
		return 0, errors.New("invalid token type")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid token claims")
	}
	log.Println(claims)

	userIDFloat, ok := claims["userID"].(float64)
	if !ok {
		return 0, errors.New("user ID not found in token claims")
	}

	return uint(userIDFloat), nil
}

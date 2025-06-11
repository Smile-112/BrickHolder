package Secrets

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtSecret = []byte("secretpass")

func GenerateToken(userId string) (string, error) {
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
		Subject:   userId,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ValidateToken(tokenString string) (*jwt.StandardClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if claims, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Отсутствуют права доступа"})
			c.Abort()
			return
		}

		claims, err := ValidateToken(token)
		if err != nil {
			if claims != nil && err.Error() == "Истекло время авторизации" {
				newToken, err := GenerateToken(claims.Subject)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Невозможно создать новый токен"})
					c.Abort()
					return
				}

				c.JSON(http.StatusUnauthorized, gin.H{
					"error":     "Истекло время авторизации",
					"new_token": newToken,
				})
				c.Abort()
				return
			}

			c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный токен"})
			c.Abort()
			return
		}

		c.Set("customerID", claims.Subject)
		c.Set("claims", claims)
		c.Next()
	}
}

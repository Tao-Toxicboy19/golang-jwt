package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		hmacSampleSecret := []byte(os.Getenv("JWT_SECRET_KEY"))
		header := c.Request.Header.Get("Authorization")

		if header == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Token"})
			c.Abort() // สิ้นสุดการทำงานของ middleware ที่นี่
			return
		}

		// ตรวจสอบว่า Token ต้องเริ่มต้นด้วย "Bearer "
		if !strings.HasPrefix(header, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Token"})
			c.Abort() // สิ้นสุดการทำงานของ middleware ที่นี่
			return
		}

		tokenString := strings.Replace(header, "Bearer ", "", 1)

		token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if token.Method != jwt.SigningMethodHS256 {
				return nil, fmt.Errorf("invalid signing method: %v", token.Header["alg"])
			}
			return hmacSampleSecret, nil
		})
		

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("user", claims["user"])
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid Token"})
		}

		c.Next()
	}
}
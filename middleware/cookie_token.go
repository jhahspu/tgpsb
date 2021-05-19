package middleware

import (
	"FS01/auth"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func CookieToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("jwt")
		if err != nil {
			c.JSON(403, "Cookie not set")
			c.Abort()
			return
		}
		clientToken := strings.TrimSpace(cookie)
		godotenv.Load(".env")
		key := os.Getenv("KEY")
		iss := os.Getenv("ISS")
		jwtWrapper := auth.JwtWrapper{
			SecretKey: key,
			Issuer:    iss,
		}
		claims, err := jwtWrapper.ValidateToken(clientToken)
		if err != nil {
			c.JSON(401, err.Error())
			c.Abort()
			return
		}
		c.Set("email", claims.Email)
		c.Next()
	}
}

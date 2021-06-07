package controller

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func SignOut(c *gin.Context) {
	godotenv.Load(".env")
	dom := os.Getenv("DOM")
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "jwt",
		Value:    "",
		Path:     "/",
		Domain:   dom,
		MaxAge:   -1,
		Expires:  time.Unix(0, 0),
		Secure:   false,
		HttpOnly: true,
		SameSite: 3,
	})
	c.JSON(200, gin.H{
		"msg": "sign out successful",
	})
}

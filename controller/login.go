package controller

import (
	"FS01/auth"
	"FS01/database"
	"FS01/models"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// LoginPayload: login request struct
type LoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginResponse: token response
type LoginResponse struct {
	Token string `json:"token"`
}

// Login
func Login(c *gin.Context) {
	var payload LoginPayload
	var user models.User

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(400, gin.H{
			"msg": "invalid json",
		})
		c.Abort()
		return
	}
	if err := database.DBClient.Get(&user, "SELECT id, name, email, password, created_at, last_login FROM users WHERE email=$1", payload.Email); err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"msg": "error getting user in db",
		})
		c.Abort()
		return
	}
	if err := user.CheckPassword(payload.Password); err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"msg": "error checking user password",
		})
		c.Abort()
		return
	}
	godotenv.Load(".env")
	key := os.Getenv("KEY")
	iss := os.Getenv("ISS")
	dom := os.Getenv("DOM")
	jwtWrapper := auth.JwtWrapper{
		SecretKey:       key,
		Issuer:          iss,
		ExpirationHours: 1,
	}
	signedToken, err := jwtWrapper.GenerateToken(user.Email)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"msg": "error signing token",
		})
		c.Abort()
		return
	}
	// in case I want to return the token as a message response
	// otherwise is set in httpOnly cookie
	// tokenResponse := LoginResponse{
	// 	Token: signedToken,
	// }
	// name(string), value(string), maxAge(int), path(string), domain(string), secure(bool), httpOnly(bool)

	// c.SetCookie("jwt", signedToken, 1, "/", dom, false, true)

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "jwt",
		Value:    signedToken,
		Path:     "/",
		Domain:   dom,
		MaxAge:   1800,
		Secure:   false,
		HttpOnly: true,
		SameSite: 3,
	})

	// c.JSON(200, tokenResponse)
	user.Password = ""

	c.JSON(200, gin.H{
		"msg": "Sign in successful",
	})
}

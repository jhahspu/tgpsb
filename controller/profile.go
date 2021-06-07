package controller

import (
	"FS01/database"
	"FS01/models"
	"log"

	"github.com/gin-gonic/gin"
)

const profileQuery = `
SELECT id, name, email, password, created_at, last_login
FROM users
WHERE email=$1
`

// Profile: return user data
func Profile(c *gin.Context) {
	var user models.User
	email, _ := c.Get("email")
	if err := database.DBClient.Get(&user, profileQuery, email); err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"msg": "error getting user in db",
		})
		c.Abort()
		return
	}
	user.Password = ""
	c.JSON(200, user)
}

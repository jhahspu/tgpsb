package controller

import (
	"FS01/models"
	"log"

	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"msg": "Invalid JSON",
		})
		c.Abort()
		return
	}
	if err := user.HashPassword(user.Password); err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"msg": "Error hashing password",
		})
		c.Abort()
		return
	}
	if err := user.CreateUser(); err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"msg": "Error creating user",
		})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{
		"msg": "Success! Now return to login page",
	})
}

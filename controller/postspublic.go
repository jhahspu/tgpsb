package controller

import (
	"FS01/database"
	"FS01/models"
	"log"

	"github.com/gin-gonic/gin"
)

// PostsFront - return 5 posts for fronpage
func PostsFront(c *gin.Context) {
	var posts []models.Post

	if err := database.DBClient.Select(&posts, "SELECT id, title, intro, stack, content, user_name, updated_at FROM posts ORDER BY updated_at DESC LIMIT 5"); err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"msg": "error getting posts from db",
		})
		c.Abort()
		return
	}

	c.JSON(200, posts)
}

func PostsByStack(c *gin.Context) {
	var posts []models.Post
	stack := c.Param("stack")
	if err := database.DBClient.Select(&posts, "SELECT id, title, intro, stack, content, user_name, updated_at FROM posts WHERE stack = $1 ORDER BY updated_at DESC", stack); err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"msg": "error getting posts from db",
		})
		c.Abort()
		return
	}

	c.JSON(200, posts)
}

// OnePost - return one post by id
func OnePost(c *gin.Context) {
	id := c.Param("id")
	var post models.Post
	if err := database.DBClient.Get(&post, "SELECT id, title, intro, stack, content, user_name, updated_at FROM posts WHERE id = $1", id); err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"msg": "error getting posts from db",
		})
		c.Abort()
		return
	}

	c.JSON(200, post)
}

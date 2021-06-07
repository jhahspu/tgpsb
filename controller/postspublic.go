package controller

import (
	"FS01/database"
	"FS01/models"
	"log"

	"github.com/gin-gonic/gin"
)

const postsFrontQuery = `
SELECT id, title, intro, stack, content, user_name, updated_at
FROM posts
ORDER BY updated_at
DESC LIMIT 5
`

// PostsFront - return 5 posts for fronpage
func PostsFront(c *gin.Context) {
	var posts []models.Post
	if err := database.DBClient.Select(&posts, postsFrontQuery); err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"msg": "error getting posts from db",
		})
		c.Abort()
		return
	}
	c.JSON(200, posts)
}

const postsByStack = `
SELECT id, title, intro, stack, content, user_name, updated_at
FROM posts
WHERE stack = $1
ORDER BY updated_at DESC
`

// PostsByStack returns all posts for stack
func PostsByStack(c *gin.Context) {
	var posts []models.Post
	stack := c.Param("stack")
	if err := database.DBClient.Select(&posts, postsByStack, stack); err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"msg": "error getting posts from db",
		})
		c.Abort()
		return
	}
	c.JSON(200, posts)
}

const onePostQuery = `
SELECT id, title, intro, stack, content, user_name, updated_at
FROM posts
WHERE id = $1
`

// OnePost - return one post by id
func OnePost(c *gin.Context) {
	id := c.Param("id")
	var post models.Post
	if err := database.DBClient.Get(&post, onePostQuery, id); err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"msg": "error getting posts from db",
		})
		c.Abort()
		return
	}
	c.JSON(200, post)
}

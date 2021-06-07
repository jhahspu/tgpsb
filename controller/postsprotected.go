package controller

import (
	"FS01/database"
	"FS01/models"
	"log"

	"github.com/gin-gonic/gin"
)

const getDBUserQuery = `
SELECT id, name, email
FROM users
WHERE email=$1
`

const createPostQuery = `
INSERT INTO posts (id, title, intro, stack, content, user_id, user_name)
VALUES (uuid_generate_v4(), $1, $2, $3, $4, $5, $6)
`

// CreatePost - create on post
func CreatePost(c *gin.Context) {
	var user models.User
	var payload models.PostPayload

	// 01: check if payload is valid json
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(400, gin.H{
			"msg": "invalid json",
		})
		c.Abort()
		return
	}

	// 02: check user and get user id and user name
	// since it's a protected route and user data is set
	// I can get the user data from gin context
	email, _ := c.Get("email")
	if err := database.DBClient.Get(&user, getDBUserQuery, email); err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"msg": "error getting user in db",
		})
		c.Abort()
		return
	}

	// 03: insert data
	// INSERT INTO posts (id, title, intro, stack, content, user_id, user_name) VALUES (uuid_generate_v4(), 'post1', 'some short description', 'stack', '# post title', 'user_id', 'user_name');
	_, err := database.DBClient.Exec(createPostQuery, &payload.Title, &payload.Intro, &payload.Stack, &payload.Content, &user.ID, &user.Name)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"msg": "error adding post, check log",
		})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{
		"msg": "success",
	})
}

const getUserPostsQuery = `
SELECT id, title, intro, stack, content, user_name, updated_at
FROM posts
WHERE user_id = $1
ORDER BY updated_at DESC
`

// GetPostForUser
func GetPostsForUser(c *gin.Context) {
	var user models.User

	// 01: check user
	email, _ := c.Get("email")
	if err := database.DBClient.Get(&user, getDBUserQuery, email); err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"msg": "error getting user in db",
		})
		c.Abort()
		return
	}

	var posts []models.Post
	// 02: get post by post_id and user_id
	// protected route
	// select * from posts where (id = 'd269e6b7-110a-45fa-b68f-48472a4acb7a' and user_id = '389e964d-ecfa-4883-a9e7-0da11db5f34c');
	if err := database.DBClient.Select(&posts, getUserPostsQuery, &user.ID); err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"msg": "error finding posts for user",
		})
		c.Abort()
		return
	}

	c.JSON(200, posts)
}

const getPostForUserQuery = `
SELECT id, title, intro, stack, content, user_name, updated_at
FROM posts
WHERE (id = $1 AND user_id = $2)
`

// GetPostForUser
func GetPostForUser(c *gin.Context) {
	var user models.User
	id := c.Param("id") // post id

	// 01: check user
	email, _ := c.Get("email")
	if err := database.DBClient.Get(&user, getDBUserQuery, email); err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"msg": "error getting user in db",
		})
		c.Abort()
		return
	}

	var post models.Post
	// 02: get post by post_id and user_id
	// protected route
	// select * from posts where (id = 'd269e6b7-110a-45fa-b68f-48472a4acb7a' and user_id = '389e964d-ecfa-4883-a9e7-0da11db5f34c');
	if err := database.DBClient.Get(&post, getPostForUserQuery, id, &user.ID); err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"msg": "error finding post for user",
		})
		c.Abort()
		return
	}

	c.JSON(200, post)
}

const deletePostQuery = `
DELETE FROM posts
WHERE id = $1
AND user_id = $2
`

// DeletePost
func DeletePost(c *gin.Context) {
	var user models.User
	id := c.Param("id")

	// 01: check user
	email, _ := c.Get("email")
	if err := database.DBClient.Get(&user, getDBUserQuery, email); err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"msg": "error getting user in db",
		})
		c.Abort()
		return
	}

	// 02: delete post by post_id and user_id
	// protected route
	_, err := database.DBClient.Exec(deletePostQuery, id, &user.ID)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"msg": "error deleting post",
		})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{
		"msg": "post deleted",
	})

}

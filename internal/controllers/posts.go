package controllers

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rujool11/chirp-core-service/internal/db"
)

func FetchAllPosts(c *gin.Context) {

}

func CreatePost(c *gin.Context) {
	userIDVal, exists := c.Get("user_id")
	if !exists {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	userId := userIDVal.(int)

	var input struct {
		Content string `json:"content" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "Failed to bind to JSON"})
		return
	}

	// get NOW() created_at tie from database itself
	createdAt := time.Now()
	query := `INSERT INTO posts (user_id, content, created_at) VALUES ($1, $2, $3)
			RETURNING id`

	var postId int
	err := db.DB.QueryRow(c, query, userId, input.Content, createdAt).Scan(&postId)
	if err != nil {
		c.JSON(500, gin.H{"error": "Could not add post in database"})
		return
	}

	c.JSON(200, gin.H{"message": "Created post",
		"post_id":    postId,
		"content":    input.Content,
		"created_at": createdAt})

}

func GetPostById(c *gin.Context) {

}

func LikePost(c *gin.Context) {

}

func UnlikePost(c *gin.Context) {

}

func DeleteOwnPost(c *gin.Context) {

}

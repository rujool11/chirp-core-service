package controllers

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rujool11/chirp-core-service/internal/db"
	"github.com/rujool11/chirp-core-service/internal/models"
)

func FetchAllPosts(c *gin.Context) {
	query := `SELECT id, user_id, content, likes_count, comments_count, created_at
			FROM posts 
			ORDER BY created_at DESC`

	rows, err := db.DB.Query(c, query)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch all posts"})
		return
	}

	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err := rows.Scan(
			&post.ID,
			&post.UserID,
			&post.Content,
			&post.LikesCount,
			&post.CommentsCount,
			&post.CreatedAt,
		)
		if err != nil {
			continue // skip problematic rows
		}

		posts = append(posts, post)
	}

	c.JSON(200, gin.H{"posts": posts})
}

func FetchPostByUser(c *gin.Context) {
	user_id := c.Param("id")
	query := `SELECT id, user_id, content, likes_count, comments_count, created_at
			FROM posts
			WHERE user_id=$1 
			ORDER BY created_at DESC`

	rows, err := db.DB.Query(c, query, user_id)
	if err != nil {
		c.JSON(500, "Failed to fetch all posts")
		return
	}

	defer rows.Close()
	var posts []models.Post

	for rows.Next() {
		var post models.Post
		err := rows.Scan(
			&post.ID,
			&post.UserID,
			&post.Content,
			&post.LikesCount,
			&post.CommentsCount,
			&post.CreatedAt,
		)
		if err != nil {
			continue
		}

		posts = append(posts, post)
	}

	c.JSON(200, gin.H{"posts": posts})

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
			RETURNING id;`

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
	id := c.Param("id")
	var post models.Post

	query := `SELECT id, user_id, content, likes_count, comments_count, created_at
			FROM posts WHERE id=$1;`

	err := db.DB.QueryRow(c, query, id).Scan(
		&post.ID,
		&post.UserID,
		&post.Content,
		&post.LikesCount,
		&post.CommentsCount,
		&post.CreatedAt,
	)
	if err != nil {
		c.JSON(500, gin.H{"error": "Could not fetch post"})
		return
	}

	c.JSON(200, gin.H{"post": post})
}

func LikePost(c *gin.Context) {

}

func UnlikePost(c *gin.Context) {

}

func DeleteOwnPost(c *gin.Context) {
	userIDVal, exists := c.Get("user_id")
	if !exists {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	postIDStr := c.Param("id")
	// .Get returns any, so we can directly use .int() to convert
	// .Param returns string, so we have to use strconv
	userID := userIDVal.(int)
	postID, err := strconv.Atoi(postIDStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid post id"})
		return
	}

	query := `DELETE FROM posts WHERE id=$1 AND user_id=$2`
	result, err := db.DB.Exec(c, query, postID, userID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete post"})
		return
	}

	if result.RowsAffected() == 0 {
		c.JSON(404, gin.H{"error": "Post not found or not your post"})
		return
	}

	c.JSON(200, gin.H{"message": "Post deleted successfully"})
}

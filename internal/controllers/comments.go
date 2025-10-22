package controllers

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rujool11/chirp-core-service/internal/db"
	"github.com/rujool11/chirp-core-service/internal/models"
)

// fetch all comments for a specific post
func FetchCommentsByPost(c *gin.Context) {
	postIDStr := c.Param("post_id")
	postID, err := strconv.Atoi(postIDStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid post id"})
		return
	}

	query := `SELECT id, post_id, user_id, content, likes_count, created_at
			FROM comments
			WHERE post_id=$1
			ORDER BY created_at DESC`

	rows, err := db.DB.Query(c, query, postID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch comments"})
		return
	}
	defer rows.Close()

	var comments []models.Comment
	for rows.Next() {
		var comment models.Comment
		err := rows.Scan(
			&comment.ID,
			&comment.PostID,
			&comment.UserID,
			&comment.Content,
			&comment.LikesCount,
			&comment.CreatedAt,
		)
		if err != nil {
			continue // skip problematic rows
		}
		comments = append(comments, comment)
	}

	c.JSON(200, gin.H{"comments": comments})
}

// create a new comment on a post
func CreateComment(c *gin.Context) {
	userIDVal, exists := c.Get("user_id")
	if !exists {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	postIDStr := c.Param("post_id")
	postID, err := strconv.Atoi(postIDStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid post id"})
		return
	}

	userID := userIDVal.(int)

	var input struct {
		Content string `json:"content" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "Failed to bind to JSON"})
		return
	}

	createdAt := time.Now()
	query := `INSERT INTO comments (post_id, user_id, content, created_at)
			VALUES ($1, $2, $3, $4)
			RETURNING id;`

	var commentID int
	err = db.DB.QueryRow(c, query, postID, userID, input.Content, createdAt).Scan(&commentID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Could not create comment"})
		return
	}

	// increment comment count in posts table
	query = `UPDATE posts SET comments_count = comments_count + 1 WHERE id=$1`
	_, err = db.DB.Exec(c, query, postID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Could not increment comment count in posts table"})
		return
	}

	c.JSON(200, gin.H{
		"message":    "Comment created",
		"comment_id": commentID,
		"content":    input.Content,
		"created_at": createdAt,
	})
}

// delete user's own comment
func DeleteOwnComment(c *gin.Context) {
	userIDVal, exists := c.Get("user_id")
	if !exists {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	commentIDStr := c.Param("id")
	userID := userIDVal.(int)
	commentID, err := strconv.Atoi(commentIDStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid comment id"})
		return
	}

	var postID int
	err = db.DB.QueryRow(c, `SELECT post_id FROM comments WHERE id=$1 AND user_id=$2`, commentID, userID).Scan(&postID)
	if err != nil {
		c.JSON(404, gin.H{"error": "Comment not found or not yours"})
		return
	}

	_, err = db.DB.Exec(c, `DELETE FROM comments WHERE id=$1 AND user_id=$2`, commentID, userID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete comment"})
		return
	}

	_, err = db.DB.Exec(c, `UPDATE posts SET comments_count = GREATEST(comments_count - 1, 0) WHERE id=$1`, postID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Could not decrement comment count in posts table"})
		return
	}

	c.JSON(200, gin.H{"message": "Comment deleted successfully"})
}

// like a comment
func LikeComment(c *gin.Context) {
	userIDVal, exists := c.Get("user_id")
	if !exists {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}
	commentIDStr := c.Param("id")
	userID := userIDVal.(int)
	commentID, err := strconv.Atoi(commentIDStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid comment id"})
		return
	}

	createdAt := time.Now()
	query := `INSERT INTO comment_likes (user_id, comment_id, created_at)
			VALUES ($1, $2, $3)
			ON CONFLICT DO NOTHING;`

	result, err := db.DB.Exec(c, query, userID, commentID, createdAt)
	if err != nil {
		c.JSON(500, gin.H{"error": "Could not add like"})
		return
	}
	if result.RowsAffected() == 0 {
		c.JSON(200, gin.H{"message": "Comment already liked"})
		return
	}

	query = `UPDATE comments SET likes_count = likes_count + 1 WHERE id=$1`
	result, err = db.DB.Exec(c, query, commentID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Could not increment like count in comments table"})
		return
	}
	if result.RowsAffected() == 0 {
		c.JSON(404, gin.H{"error": "Comment not found"})
		return
	}

	c.JSON(200, gin.H{"message": "Comment liked"})
}

// unlike comment
func UnlikeComment(c *gin.Context) {
	userIDVal, exists := c.Get("user_id")
	if !exists {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}
	commentIDStr := c.Param("id")
	userID := userIDVal.(int)
	commentID, err := strconv.Atoi(commentIDStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid comment id"})
		return
	}

	query := `DELETE FROM comment_likes WHERE user_id=$1 AND comment_id=$2`
	result, err := db.DB.Exec(c, query, userID, commentID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Could not remove like"})
		return
	}
	if result.RowsAffected() == 0 {
		c.JSON(200, gin.H{"message": "Comment not liked"})
		return
	}

	query = `UPDATE comments SET likes_count = GREATEST(likes_count - 1, 0) WHERE id=$1`
	result, err = db.DB.Exec(c, query, commentID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Could not decrement like count in comments table"})
		return
	}
	if result.RowsAffected() == 0 {
		c.JSON(404, gin.H{"error": "Comment not found"})
		return
	}

	c.JSON(200, gin.H{"message": "Comment unliked"})
}

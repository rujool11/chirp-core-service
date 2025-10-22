package db

import (
	"context"
	"log"
)

func CreateCommentLikeTableIfDoesNotExist() {
	query := `
	CREATE TABLE IF NOT EXISTS comment_likes (
		user_id INT NOT NULL,
		comment_id INT NOT NULL,
		created_at TIMESTAMP DEFAULT NOW(),
		PRIMARY KEY(user_id, comment_id),
		FOREIGN KEY (comment_id) REFERENCES comments(id) ON DELETE CASCADE
	);
	`
	_, err := DB.Exec(context.Background(), query)
	if err != nil {
		log.Fatalf("SQL query error when creating comment_likes table: %v", err)
	}
	log.Println("CommentLikes table is ready")
}

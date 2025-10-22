package db

import (
	"context"
	"log"
)

func CreatePostLikeTableIfDoesNotExist() {
	query := `
	CREATE TABLE IF NOT EXISTS post_likes (
		user_id INT NOT NULL,
		post_id INT NOT NULL,
		created_at TIMESTAMP DEFAULT NOW(),
		PRIMARY KEY(user_id, post_id),
		FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE
	);
	`
	_, err := DB.Exec(context.Background(), query)
	if err != nil {
		log.Fatalf("SQL query error when creating post_likes table: %v", err)
	}
	log.Println("PostLikes table is ready")
}

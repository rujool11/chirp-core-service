package db

import (
	"context"
	"log"
)

func CreateCommentTableIfDoesNotExist() {
	query := `
	CREATE TABLE IF NOT EXISTS comments (
		id SERIAL PRIMARY KEY,
		post_id INT NOT NULL,
		user_id INT NOT NULL,
		content TEXT NOT NULL,
		likes_count INT DEFAULT 0,
		created_at TIMESTAMP DEFAULT NOW(),
		FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE
	);
	`
	_, err := DB.Exec(context.Background(), query)
	if err != nil {
		log.Fatalf("SQL query error when creating comments table: %v", err)
	}
	log.Println("Comments table is ready")
}

package db

import (
	"context"
	"log"
)

func CreatePostTableIfDoesNotExist() {
	query := `
	CREATE TABLE IF NOT EXISTS posts (
		id SERIAL PRIMARY KEY,
		user_id INT NOT NULL,
		content TEXT NOT NULL,
		likes_count INT DEFAULT 0,
		comments_count INT DEFAULT 0,
		created_at TIMESTAMP DEFAULT NOW()
	);
	`
	_, err := DB.Exec(context.Background(), query)
	if err != nil {
		log.Fatalf("SQL query error when creating posts table: %v", err)
	}
	log.Println("Posts table is ready")
}

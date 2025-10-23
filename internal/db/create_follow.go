package db

import (
	"context"
	"log"
)

func CreateFollowTableIfDoesNotExist() {
	query := `
	CREATE TABLE IF NOT EXISTS follow (
		follower_id INT NOT NULL,
		following_id INT NOT NULL,
		created_at TIMESTAMP DEFAULT NOW(),
		PRIMARY KEY(follower_id, following_id),
		FOREIGN KEY (follower_id) REFERENCES users(id) ON DELETE CASCADE,
		FOREIGN KEY (following_id) REFERENCES users(id) ON DELETE CASCADE
	);
	`
	_, err := DB.Exec(context.Background(), query)
	if err != nil {
		log.Fatalf("SQL query error when creating follow table: %v", err)
	}
	log.Println("Follow table is ready")
}

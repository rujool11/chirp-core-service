package models

import "time"

type Post struct {
	ID            int       `json:"id"`
	UserID        int       `json:"user_id"`
	Content       string    `json:"content"`
	LikesCount    int       `json:"likes_count"`
	CommentsCount int       `json:"comments_count"`
	CreatedAt     time.Time `json:"created_at"`
}

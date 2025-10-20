package models

import "time"

type Comment struct {
	ID         int       `json:"id"`
	PostID     int       `json:"post_id"`
	UserID     int       `json:"user_id"` // user who commented
	Content    string    `json:"content"`
	LikesCount int       `json:"likes_count"`
	CreatedAt  time.Time `json:"created_at"`
}

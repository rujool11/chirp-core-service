package models

import "time"

type PostLike struct {
	UserID    int       `json:"user_id"` // user who liked the post
	PostID    int       `json:"post_id"`
	CreatedAt time.Time `json:"created_at"`
}

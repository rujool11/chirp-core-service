package models

import "time"

type CommentLike struct {
	UserID    int       `json:"user_id"` // user who made the comment
	CommentID int       `json:"comment_id"`
	CreatedAt time.Time `json:"created_at"`
}

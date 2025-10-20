package models

import "time"

type User struct {
	// struct tags used to indicate JSON key for JSON encoder/decoder
	ID             int       `json:"id"`
	Username       string    `json:"username"`
	Email          string    `json:"email"`
	PasswordHash   string    `json:"-"` // ignore this when converting to json
	Bio            string    `json:"bio"`
	LikesCount     int       `json:"likes_count"`
	FollowersCount int       `json:"followers_count"`
	FollowingCount int       `json:"following_count"`
	TweetsCount    int       `json:"tweets_count"`
	CreatedAt      time.Time `json:"created_at"`
}

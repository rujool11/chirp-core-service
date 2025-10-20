package models

import "time"

type Follow struct {
	FollowerID  int       `json:"follower_id"`  // follower
	FollowingID int       `json:"following_id"` // user being followed
	CreatedAt   time.Time `json:"created_at"`
}

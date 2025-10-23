package controllers

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rujool11/chirp-core-service/internal/db"
)

func FollowUser(c *gin.Context) {
	userIDVal, exists := c.Get("user_id")
	if !exists {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	targetIDStr := c.Param("id")
	targetID, err := strconv.Atoi(targetIDStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid id"})
		return
	}

	followerID := userIDVal.(int)
	if followerID == targetID {
		c.JSON(400, gin.H{"error": "You cannot follow yourself"})
		return
	}

	query := `INSERT INTO follow (follower_id, following_id, created_at)
			VALUES ($1, $2, $3)
			ON CONFLICT DO NOTHING`

	result, err := db.DB.Exec(c, query, followerID, targetID, time.Now())
	if err != nil {
		c.JSON(500, gin.H{"error": "Could not insert into database"})
		return
	}

	if result.RowsAffected() == 0 {
		c.JSON(200, gin.H{"message": "Already followed user"})
		return
	}

	_, err = db.DB.Exec(c, `UPDATE users SET following_count = following_count + 1 WHERE id=$1`, followerID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Could not increment following count"})
		return
	}

	_, err = db.DB.Exec(c, `UPDATE users SET followers_count = followers_count + 1 WHERE id=$1`, targetID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Could not increment followers count"})
		return
	}

	c.JSON(200, gin.H{"message": "Followed user"})
}

func UnfollowUser(c *gin.Context) {
	userIDVal, exists := c.Get("user_id")
	if !exists {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	targetIDStr := c.Param("id")
	targetID, err := strconv.Atoi(targetIDStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid id"})
		return
	}

	followerID := userIDVal.(int)
	if followerID == targetID {
		c.JSON(400, gin.H{"error": "You cannot unfollow yourself"})
		return
	}

	query := `DELETE FROM follow WHERE follower_id=$1 AND following_id=$2`
	result, err := db.DB.Exec(c, query, followerID, targetID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Could not remove follow from database"})
		return
	}

	if result.RowsAffected() == 0 {
		c.JSON(200, gin.H{"message": "You were not following this user"})
		return
	}

	_, err = db.DB.Exec(c, `UPDATE users SET following_count = GREATEST(following_count - 1, 0) WHERE id=$1`, followerID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Could not decrement following count"})
		return
	}

	_, err = db.DB.Exec(c, `UPDATE users SET followers_count = GREATEST(followers_count - 1, 0) WHERE id=$1`, targetID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Could not decrement followers count"})
		return
	}

	c.JSON(200, gin.H{"message": "Unfollowed user"})
}

func GetFollowers(c *gin.Context) {
	userIDStr := c.Param("id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid user id"})
		return
	}

	query := `SELECT u.id, u.username, u.bio, u.followers_count, u.following_count
			  FROM follow f
			  JOIN users u ON f.follower_id = u.id
			  WHERE f.following_id = $1`

	rows, err := db.DB.Query(c, query, userID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch followers"})
		return
	}
	defer rows.Close()

	type Follower struct {
		ID             int    `json:"id"`
		Username       string `json:"username"`
		Bio            string `json:"bio"`
		FollowersCount int    `json:"followers_count"`
		FollowingCount int    `json:"following_count"`
	}

	var followers []Follower
	for rows.Next() {
		var f Follower
		if err := rows.Scan(&f.ID, &f.Username, &f.Bio, &f.FollowersCount, &f.FollowingCount); err != nil {
			continue
		}
		followers = append(followers, f)
	}

	c.JSON(200, gin.H{"followers": followers})
}

func GetFollowing(c *gin.Context) {
	userIDStr := c.Param("id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid user id"})
		return
	}

	query := `SELECT u.id, u.username, u.bio, u.followers_count, u.following_count
			  FROM follow f
			  JOIN users u ON f.following_id = u.id
			  WHERE f.follower_id = $1`

	rows, err := db.DB.Query(c, query, userID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch following"})
		return
	}
	defer rows.Close()

	type Following struct {
		ID             int    `json:"id"`
		Username       string `json:"username"`
		Bio            string `json:"bio"`
		FollowersCount int    `json:"followers_count"`
		FollowingCount int    `json:"following_count"`
	}

	var following []Following
	for rows.Next() {
		var f Following
		if err := rows.Scan(&f.ID, &f.Username, &f.Bio, &f.FollowersCount, &f.FollowingCount); err != nil {
			continue
		}
		following = append(following, f)
	}

	c.JSON(200, gin.H{"following": following})
}

package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rujool11/chirp-core-service/internal/controllers"
	"github.com/rujool11/chirp-core-service/internal/db"
	"github.com/rujool11/chirp-core-service/internal/middleware"
)

func main() {
	// initialize DB connection
	db.InitDB()
	defer db.DB.Close()
	db.CreatePostTableIfDoesNotExist()
	db.CreateCommentTableIfDoesNotExist()
	db.CreatePostLikeTableIfDoesNotExist()
	db.CreateCommentLikeTableIfDoesNotExist()
	db.CreateFollowTableIfDoesNotExist()

	err := godotenv.Load()
	if err != nil {
		log.Println(".env not found")
	}

	r := gin.Default()

	post := r.Group("/posts")
	{
		post.GET("/", controllers.FetchAllPosts)
		post.POST("/", middleware.AuthMiddleware(), controllers.CreatePost)
		post.GET("/user/:id", controllers.FetchPostByUser)
		post.GET("/:id", controllers.GetPostById)
		post.DELETE("/:id", middleware.AuthMiddleware(), controllers.DeleteOwnPost)
		post.POST("/:id/like", middleware.AuthMiddleware(), controllers.LikePost)
		post.DELETE("/:id/unlike", middleware.AuthMiddleware(), controllers.UnlikePost)
	}

	comment := r.Group("/comments")
	{
		comment.GET("/post/:post_id", controllers.FetchCommentsByPost)
		comment.POST("/post/:post_id", middleware.AuthMiddleware(), controllers.CreateComment)
		comment.DELETE("/:id", middleware.AuthMiddleware(), controllers.DeleteOwnComment)
		comment.POST("/:id/like", middleware.AuthMiddleware(), controllers.LikeComment)
		comment.DELETE("/:id/unlike", middleware.AuthMiddleware(), controllers.UnlikeComment)
	}

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello from chirp-core-service",
		})
	})

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8002"
		log.Println("Defaulting PORT to 8002")
	}

	r.Run(":" + PORT)
}

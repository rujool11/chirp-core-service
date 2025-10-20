package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rujool11/chirp-core-service/internal/db"
)

func main() {
	// initialize DB connection
	db.InitDB()
	defer db.DB.Close()

	err := godotenv.Load()
	if err != nil {
		log.Println(".env not found")
	}

	r := gin.Default()

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

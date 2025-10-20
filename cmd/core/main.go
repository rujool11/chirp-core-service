package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

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

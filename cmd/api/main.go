package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/fahmi/go-order-api/internal/config"
)

func main() {
	_ = godotenv.Load()

	db, err := config.InitDB()
	if err != nil {
		log.Println("warning: database not connected, running without DB")
	}

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		status := "ok"
		if db == nil {
			status = "db_disconnected"
		}
		c.JSON(200, gin.H{"status": status})
	})

	log.Println("Server running on :8080")
	r.Run(":8080")
}

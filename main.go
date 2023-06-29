package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	app := gin.New()
	app.Handle("GET", "ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	app.Run(":8080")
}

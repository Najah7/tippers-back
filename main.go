package main

import (
	"log"
	"tippers-back/db"
	"tippers-back/middleware"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("読み込み出来ませんでした: %v", err)
	}
	r := gin.Default()
	db.NewDB()
	r.Use(middleware.AuthMiddleware())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}

package main

import (
	"log"
	"tippers-back/controller"
	"tippers-back/db"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("読み込み出来ませんでした: %v", err)
	}

	db.NewDB()
	r := controller.Router()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}

package main

import (
	"log"
	"tippers-back/controller"
	"tippers-back/db"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("読み込み出来ませんでした: %v", err)
	}

	db.NewDB()
	r := controller.Router()
	r.Run()
}

package controller

import (
	"log"
	"tippers-back/db"
)

type handler struct {
	db *db.DB
}

func (h *handler) Init() {
	var err error
	if h.db, err = db.NewDB(); err != nil {
		log.Fatalf("Database connection failed")
	}
	if err=h.db.CreateTable();err!=nil{
		log.Fatalf("table failed")
	}
}

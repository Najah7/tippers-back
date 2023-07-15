package controller

import (
	"log"
	"net/http"
	"tippers-back/db"
	"tippers-back/db/model"

	"github.com/gin-gonic/gin"
)

type handler struct {
	db *db.DB
}

func (h *handler) Init() {
	var err error
	if h.db, err = db.NewDB(); err != nil {
		log.Fatalf("Database connection failed")
	}
	if err = h.db.CreateTable(); err != nil {
		log.Fatalf("table failed")
	}
}

func (h *handler) RegisterUser(c *gin.Context) {
	var user model.User
	var err error
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err = h.db.RegisterUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

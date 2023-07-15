package controller

import (
	"log"
	"net/http"
	"strconv"
	"tippers-back/db"
	"tippers-back/db/table"
	"tippers-back/lib"
	"tippers-back/service"

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

func (h *handler) GetUsers(c *gin.Context) {
	var users *[]table.User
	var err error
	users, err = h.db.GetUsers(users)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (h *handler) GetUser(c *gin.Context) {
	var user *table.User

	stringID := c.Param("id")
	id, err := strconv.Atoi(stringID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err = h.db.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *handler) RegisterUser(c *gin.Context) {
	var user *table.User
	var err error
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.Password = service.HashPassword(user.Password)
	user, err = h.db.RegisterUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *handler) UpdateUser(c *gin.Context) {
	stringID := c.Param("id")
	id, err := strconv.Atoi(stringID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	float64UserID := c.MustGet("user_id").(float64)
	UserID := int(float64UserID)
	if id != UserID {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Authorization header is missing",
		})
		return
	}
	var user *table.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dbUser, err := h.db.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dbUser = lib.UpdateUser(user, dbUser)

	user, err = h.db.UpdateUser(dbUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *handler) Login(c *gin.Context) {
	type resposen struct {
		Token string `json:"token"`
	}
	var response resposen

	var user table.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	dbUser, err := h.db.GetUserByMail(user.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := service.ComparePassword(dbUser.Password, user.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	response.Token, err = service.JwtGenerate(*dbUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

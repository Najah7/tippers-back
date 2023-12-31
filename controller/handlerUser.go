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
	"github.com/google/uuid"
	imgupload "github.com/olahol/go-imageupload"
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
	users, err := h.db.GetUsers()
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
	userID := int(float64UserID)
	if id != userID {
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

	dbUser, err := h.db.GetUserByID(userID)
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

func (h *handler) DeleteUser(c *gin.Context) {
	stringID := c.Param("id")
	id, err := strconv.Atoi(stringID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	float64UserID := c.MustGet("user_id").(float64)
	userID := int(float64UserID)
	if id != userID {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Authorization header is missing",
		})
		return
	}

	err = h.db.DeleteUserByID(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, true)
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

func (h *handler) UploadUserProfile(c *gin.Context) {
	type response struct {
		ProfileImageURL string `json:"profileImageURL"`
	}

	img, err := imgupload.Process(c.Request, "file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	webp, err := lib.ConvertToWebp(img)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	uuid, err := uuid.NewRandom()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	filename := uuid.String() + ".webp"
	filepass, err := lib.UploadImage(webp, filename, "user")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	float64UserID := c.MustGet("user_id").(float64)
	userID := int(float64UserID)
	if err := h.db.UpdateProfileImageURLIDByID(userID, filepass); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, &response{
		ProfileImageURL: filepass,
	})
}

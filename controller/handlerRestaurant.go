package controller

import (
	"net/http"
	"strconv"
	"tippers-back/db/table"
	"tippers-back/lib"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	imgupload "github.com/olahol/go-imageupload"
)

func (h *handler) GetRestaurants(c *gin.Context) {
	restaurants, err := h.db.GetRestaurants()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, restaurants)
}

func (h *handler) GetRestaurant(c *gin.Context) {
	var restaurant *table.Restaurant
	stringID := c.Param("id")
	id, err := strconv.Atoi(stringID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	restaurant, err = h.db.GetRestaurantByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, restaurant)
}

func (h *handler) RegisterRestaurant(c *gin.Context) {
	var restaurant *table.Restaurant
	var err error
	if err := c.ShouldBindJSON(&restaurant); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	restaurant, err = h.db.RegisterRestaurant(restaurant)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	float64UserID := c.MustGet("user_id").(float64)
	err = h.db.UpdateUserRestaurantIDByID(int(float64UserID), int(restaurant.ID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, restaurant)
}

func (h *handler) UploadRestaurantProfile(c *gin.Context) {
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
	filepass, err := lib.UploadImage(webp, filename, "restaurant")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	float64UserID := c.MustGet("user_id").(float64)
	userID := int(float64UserID)
	user, err := h.db.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.db.UpdateRestaurantProfileImageURLIDByID(user.RestaurantID, filepass); err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, &response{
		ProfileImageURL: filepass,
	})
}

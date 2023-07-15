package controller

import (
	"net/http"
	"strconv"
	"tippers-back/db/table"

	"github.com/gin-gonic/gin"
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

package controller

import (
	"net/http"

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

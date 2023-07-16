package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *handler) GetStaffs(c *gin.Context) {
	stringID := c.Param("restaurant-id")
	id, err := strconv.Atoi(stringID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tips, err := h.db.GetStaffByRestaurantID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tips)
}

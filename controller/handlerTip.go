package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) GetTips(c *gin.Context) {
	float64UserID := c.MustGet("user_id").(float64)
	userID := int(float64UserID)
	tips, err := h.db.GetTipsBySenderID(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tips)
}

package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) GetEmployedUsers(c *gin.Context) {
	users, err := h.db.GetEmployedUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

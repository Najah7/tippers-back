package controller

import (
	"net/http"
	"tippers-back/db/table"

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

func (h *handler) GetTotalAmount(c *gin.Context) {
	type resposen struct {
		Total int `json:"total"`
	}
	var totalTip resposen
	float64UserID := c.MustGet("user_id").(float64)
	userID := int(float64UserID)
	tipTotalDB, err := h.db.GetTipAmountBySenderID(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	totalTip.Total = *tipTotalDB

	c.JSON(http.StatusOK, totalTip)
}

func (h *handler) SendTip(c *gin.Context) {
	var tip *table.Tip
	var err error
	if err := c.ShouldBindJSON(&tip); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	float64UserID := c.MustGet("user_id").(float64)
	tip.SenderID = int(float64UserID)
	tip, err = h.db.SendTip(tip)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.db.UpdateUserMoneyByID(tip.SenderID, tip.ReceiverID, tip.Amount); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tip)
}

package handlers

import (
	"net/http"

	"antin0.de/comm-relay/models"
	"github.com/gin-gonic/gin"
)

type CreateEmailTargetRequestBody struct {
	ChannelID    uint   `json:"channelId"`
	EmailAddress string `json:"emailAddress"`
}

// POST /createEmailTarget
func (h *HandlerParams) CreateEmailTarget() gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestBody CreateEmailTargetRequestBody
		if err := c.BindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid request",
			})
			return
		}
		if requestBody.EmailAddress == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "You must provide an email address",
			})
			return
		}

		// Try to find the channel with provided ID
		result := h.Db.Find(&models.Channel{}).Where("id = ?", requestBody.ChannelID)
		if result.RowsAffected == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid channel ID",
			})
			return
		}

		target := models.EmailTarget{ChannelID: requestBody.ChannelID, EmailAddress: requestBody.EmailAddress}
		result = h.Db.Create(&target)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to create target",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"id": target.ID,
		})
	}
}

// POST /updateEmailTarget
func (h *HandlerParams) UpdateEmailTarget() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"pong": true,
		})
	}
}

// POST /deleteEmailTarget
func (h *HandlerParams) DeleteEmailTarget() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"pong": true,
		})
	}
}

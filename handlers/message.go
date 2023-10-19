package handlers

import (
	"net/http"

	"antin0.de/comm-relay/models"
	"github.com/gin-gonic/gin"
)

type SendMessageRequestBody struct {
	ChannelID int    `json:"channelId"`
	Title     string `json:"title"`
	Content   string `json:"content"`
}

// POST /sendMessage
func (h *HandlerParams) SendMessage() gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestBody SendMessageRequestBody
		if err := c.BindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid request",
			})
			return
		}

		if requestBody.Title == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "You must provide a title",
			})
			return
		}

		if requestBody.Content == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "You must provide content",
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

		message := models.Message{
			Title:     requestBody.Title,
			Content:   requestBody.Content,
			ChannelID: requestBody.ChannelID,
		}
		result = h.Db.Create(&message)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to send message",
			})
			return
		}

		// TODO: actually process the messages

		c.JSON(http.StatusOK, gin.H{
			"id": message.ID,
		})
	}
}

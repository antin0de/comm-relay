package handlers

import (
	"net/http"
	"net/smtp"
	"os"

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
		emailTargets := []models.EmailTarget{}
		h.Db.Where("channel_id = ?", requestBody.ChannelID).Find(&emailTargets)
		for _, target := range emailTargets {
			smtpAuth := smtp.PlainAuth(
				"",
				os.Getenv("SMTP_USER"),
				os.Getenv("SMTP_PASSWORD"),
				os.Getenv("SMTP_HOST"),
			)
			to := []string{target.EmailAddress}
			msg := []byte("To: " + target.EmailAddress + "\r\n" +
				"Subject: " + requestBody.Title + "\r\n" +
				"\r\n" +
				requestBody.Content + "\r\n",
			)
			err := smtp.SendMail(
				os.Getenv("SMTP_HOST")+":"+os.Getenv("SMTP_PORT"),
				smtpAuth,
				os.Getenv("SMTP_USER"),
				to, msg,
			)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "Failed to send message",
				})
				println(err.Error())
				return
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"id": message.ID,
		})
	}
}

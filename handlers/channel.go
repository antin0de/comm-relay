package handlers

import (
	"net/http"

	"antin0.de/comm-relay/models"
	"github.com/gin-gonic/gin"
)

type CreateChannelRequestBody struct {
	Name string `json:"name"`
}

// POST /createChannel
func (h *HandlerParams) CreateChannel() gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestBody CreateChannelRequestBody
		if err := c.BindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid request",
			})
			return
		}
		if requestBody.Name == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "You must provide a name",
			})
			return
		}

		channel := models.Channel{Name: requestBody.Name}
		result := h.Db.Create(&channel)

		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Could not create channel",
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"id": channel.ID,
		})
	}
}

type UpdateChannelRequestBody struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// POST /updateChannel
func (h *HandlerParams) UpdateChannel() gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestBody UpdateChannelRequestBody
		if err := c.BindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid request",
			})
			return
		}
		if requestBody.Name == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "You must provide a name",
			})
			return
		}

		result := h.Db.Model(&models.Channel{}).Where("id = ?", requestBody.ID).Update("name", requestBody.Name)

		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Could not update channel",
			})
		}

		if result.RowsAffected == 0 {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Channel not found",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"id": requestBody.ID,
		})
	}
}

// POST /listChannels
func (h *HandlerParams) ListChannels() gin.HandlerFunc {
	return func(c *gin.Context) {
		var channels []models.Channel
		h.Db.Find(&channels)

		var response []gin.H
		for _, channel := range channels {
			var emailTargets []gin.H
			emailTargetModels := []models.EmailTarget{}
			h.Db.Where("channel_id = ?", channel.ID).Find(&emailTargetModels)
			for _, emailTargetModel := range emailTargetModels {
				emailTargets = append(emailTargets, gin.H{
					"id":           emailTargetModel.ID,
					"emailAddress": emailTargetModel.EmailAddress,
				})
			}
			response = append(response, gin.H{
				"id":           channel.ID,
				"name":         channel.Name,
				"emailTargets": emailTargets,
				"createdAt":    channel.CreatedAt,
				"updatedAt":    channel.UpdatedAt,
			})
		}
		c.JSON(http.StatusOK, response)
	}
}

type DeleteChannelRequestBody struct {
	ID int `json:"id"`
}

// POST /deleteChannel
func (h *HandlerParams) DeleteChannel() gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestBody DeleteChannelRequestBody
		if err := c.BindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid request",
			})
			return
		}
		// delete a channel model with ID
		result := h.Db.Delete(&models.Channel{}, requestBody.ID)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Could not delete channel",
			})
		}
		if result.RowsAffected == 0 {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Channel not found",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"id": requestBody.ID,
		})
	}
}

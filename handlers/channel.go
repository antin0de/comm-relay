package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// POST /createChannel
func (h *HandlerParams) CreateChannel() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"pong": true,
		})
	}
}

// POST /updateChannel
func (h *HandlerParams) UpdateChannel() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"pong": true,
		})
	}
}

// POST /listChannels
func (h *HandlerParams) ListChannels() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"pong": true,
		})
	}
}

// POST /deleteChannel
func (h *HandlerParams) DeleteChannel() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"pong": true,
		})
	}
}

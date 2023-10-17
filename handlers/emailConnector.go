package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// POST /createEmailConnector
func (h *HandlerParams) CreateEmailConnector() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"pong": true,
		})
	}
}

// POST /updateEmailConnector
func (h *HandlerParams) UpdateEmailConnector() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"pong": true,
		})
	}
}

// POST /deleteEmailConnector
func (h *HandlerParams) DeleteEmailConnector() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"pong": true,
		})
	}
}

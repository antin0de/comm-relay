package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// POST /createEmailTarget
func (h *HandlerParams) CreateEmailTarget() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"pong": true,
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

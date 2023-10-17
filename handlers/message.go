package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// POST /sendMessage
func (h *HandlerParams) SendMessage() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"pong": true,
		})
	}
}

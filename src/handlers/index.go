package handlers

import "github.com/gin-gonic/gin"

// Index is a handler for a ping test
func Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "ok",
	})
}

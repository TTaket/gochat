package service

import "github.com/gin-gonic/gin"

// GetPing
// @Description ping test
// @Tags Ping
// @Accept json
// @Produce json
// @Success 200 {string} pong
// @Router /ping [get]
func GetPing(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

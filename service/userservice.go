package service

import (
	"github.com/TTaket/gochat/models"
	"github.com/TTaket/gochat/utils"
	"github.com/gin-gonic/gin"
)

func GetUserList(c *gin.Context) {
	var userList []*models.UserBasic
	var ok error
	if userList, ok = models.GetUserList(utils.GetDB()); ok != nil {
		c.JSON(200, gin.H{
			"message": "server error",
			"error":   ok.Error(),
		})
		return
	}

	var s []string
	for _, user := range userList {
		s = append(s, user.Name)
	}
	c.JSON(200, gin.H{
		"message":   "user list",
		"user_list": s,
	})
}

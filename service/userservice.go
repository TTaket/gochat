package service

import (
	"github.com/TTaket/gochat/models"
	"github.com/TTaket/gochat/utils"
	"github.com/gin-gonic/gin"
)

// GetUserList
// @Description get userlist
// @Tags UserInfo
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{} "user list是用户名列表"
// @Failure 500 {object} map[string]interface{} "错误信息"
// @Router /user/getUserList [get]
func GetUserList(c *gin.Context) {
	var userList []*models.UserBasic
	var ok error
	if userList, ok = models.GetUserList(utils.GetDB()); ok != nil {
		c.JSON(500, gin.H{
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

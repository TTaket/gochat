package service

import (
	"strconv"

	"github.com/TTaket/gochat/models"
	"github.com/TTaket/gochat/utils"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

// GetUserList
// @Description get userlist
// @Tags Admin
// @Success 200 {object} map[string]interface{} "user list是用户列表"
// @Failure 500 {object} map[string]interface{} "错误信息"
// @Router /admin/getUserList [get]
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
		s = append(s, user.String())
		s = append(s, "\n")
	}
	c.JSON(200, gin.H{
		"message":   "user list",
		"user_list": s,
	})
}

// GetUserByID
// @Description get userinfo by id
// @Tags Admin
// @param id query int true "用户ID"
// @Success 200 {object} map[string]interface{} "user list是用户名列表"
// @Failure 500 {object} map[string]interface{} "错误信息"
// @Router /admin/getUserByID [get]
func GetUserByID(c *gin.Context) {
	var userInfo *models.UserBasic
	var ok error
	userId := c.Query("id")
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "invalid user id",
			"error":   err.Error(),
		})
		return
	}

	if userInfo, ok = models.GetUserByID(utils.GetDB(), userIdInt); ok != nil {
		c.JSON(500, gin.H{
			"message": "server error",
			"error":   ok.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message":  "user list",
		"userInfo": userInfo.String(),
	})
}

// CreateUser
// @Description create user
// @Tags Admin
// @param name query string true "用户名 6-20位"
// @param password query string true "密码 6-20位"
// @param phone query string true "手机号"
// @param email query string false "邮箱"
// @Success 200 {object} map[string]interface{} "user list是用户名列表"
// @Failure 500 {object} map[string]interface{} "错误信息"
// @Router /admin/createUser [get]
func CreateUser(c *gin.Context) {
	user := models.UserBasic{}
	user.Name = c.Query("name")
	user.PassWord = c.Query("password")
	user.Phone = c.Query("phone")
	user.Email = c.Query("email")
	user.Salt = utils.GenerateSalt()

	//检验参数
	if ok, err := govalidator.ValidateStruct(user); !ok {
		c.JSON(400, gin.H{
			"message": "invalid user info",
			"error":   err.Error(),
		})
		return
	}

	user.PassWord = utils.GeneratePassword(user.PassWord, user.Salt)

	if err := models.CreateUser(utils.GetDB(), &user); err != nil {
		c.JSON(500, gin.H{
			"message": "server error",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "create user success",
		"user":    user.String(),
	})
}

// DeleteUserByID
// @Description delete user
// @Tags Admin
// @param id query int true "用户ID"
// @Success 200 {object} map[string]interface{} "Delete user success"
// @Failure 500 {object} map[string]interface{} "错误信息"
// @Router /admin/deleteUserByID [get]
func DeleteUserByID(c *gin.Context) {
	userId := c.Query("id")
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "invalid user id",
			"error":   err.Error(),
		})
		return
	}
	if err := models.DeleteUserByID(utils.GetDB(), userIdInt); err != nil {
		c.JSON(500, gin.H{
			"message": "server error",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Delete user success",
	})
}

// UpdateUser
// @Description update user
// @Tags Admin
// @param id query int true "用户ID"
// @param name query string false "用户名"
// @param password query string false "密码"
// @param phone query string false "手机号"
// @param email query string false "邮箱"
// @param identity query string false "身份"
// Success 200 {object} map[string]interface{} "Update user success"
// @Failure 500 {object} map[string]interface{} "错误信息"
// @Router /admin/updateUser [post]
func UpdateUser(c *gin.Context) {
	userId := c.Query("id")
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "invalid user id",
			"error":   err.Error(),
		})
		return
	}
	//get old user info
	user, ok := models.GetUserByID(utils.GetDB(), userIdInt)
	if ok != nil {
		c.JSON(500, gin.H{
			"message": "server error",
			"error":   ok.Error(),
		})
		return
	}

	//update user info
	user.ID = uint(userIdInt)
	if name := c.Query("name"); name != "" {
		user.Name = name
	}
	if PassWord := c.Query("password"); PassWord != "" {
		user.PassWord = c.Query("password")
	}
	if phone := c.Query("phone"); phone != "" {
		user.Phone = c.Query("phone")
	}
	if email := c.Query("email"); email != "" {
		user.Email = c.Query("email")
	}
	if identity := c.Query("identity"); identity != "" {
		user.Identity = c.Query("identity")
	}

	//检验参数
	if ok, err := govalidator.ValidateStruct(user); !ok {
		c.JSON(400, gin.H{
			"message": "invalid user info",
			"error":   err.Error(),
		})
		return
	}

	user.PassWord = utils.GeneratePassword(user.PassWord, user.Salt)

	if err := models.UpdateUser(utils.GetDB(), user); err != nil {
		c.JSON(500, gin.H{
			"message": "server error",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Update user success",
	})
}

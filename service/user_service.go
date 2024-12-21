package service

import (
	"github.com/TTaket/gochat/models"
	"github.com/TTaket/gochat/utils"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

// Login
// @Description user login
// @Tags User
// @param name query string true "用户名"
// @param password query string true "密码"
// @Success 200 {object} map[string]interface{} "Login success"
// @Failure 500 {object} map[string]interface{} "错误信息"
// @Router /login [post]
func Login(c *gin.Context) {
	user := models.UserBasic{}
	user.Name = c.Query("name")
	user.PassWord = c.Query("password")

	//检验参数
	if ok, err := govalidator.ValidateStruct(user); !ok {
		c.JSON(400, gin.H{
			"message": "invalid user info",
			"error":   err.Error(),
		})
		return
	}

	//get user info
	userInfo, ok := models.GetUserByName(utils.GetDB(), user.Name)
	if ok != nil {
		c.JSON(500, gin.H{
			"message": "server error",
			"error":   ok.Error(),
		})
		return
	}

	//check password
	if userInfo.PassWord != utils.GeneratePassword(user.PassWord, userInfo.Salt) {
		c.JSON(400, gin.H{
			"message": "password error",
		})
		return
	}

	token, _ := utils.GenerateToken(user.Name)
	redisclient := utils.GetRedisClient()
	redisclient.Set(user.Name, token, utils.GetJwtExpireDuration())
	c.JSON(200, gin.H{
		"message": "Login success",
		"token":   token,
	})
}

// register
// @Description register user
// @Tags User
// @param name query string true "用户名 6-20位"
// @param password query string true "密码 6-20位"
// @param phone query string true "手机号"
// @param email query string false "邮箱"
// @Success 200 {object} map[string]interface{} "user list是用户名列表"
// @Failure 500 {object} map[string]interface{} "错误信息"
// @Router /register [post]
func Register(c *gin.Context) {
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

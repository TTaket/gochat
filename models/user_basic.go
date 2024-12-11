package models

import (
	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model

	Name          string
	PassWord      string
	Phone         string
	Email         string
	Identity      string
	ClientIp      string
	ClientPort    string
	LoginTime     string
	LoginOutTime  string
	HeartbeatTime string
	IsLogout      bool
	DeviceInfo    string
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}
func GetUserList(DB *gorm.DB) ([]*UserBasic, error) {
	// 初始化一个切片来存储查询结果
	var users []*UserBasic

	// 执行查询
	result := DB.Find(&users)
	if result.Error != nil {
		// 如果发生错误，返回错误
		return nil, result.Error
	}

	// 查询成功，返回结果和 nil 错误
	return users, nil
}

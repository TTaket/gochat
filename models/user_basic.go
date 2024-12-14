package models

import (
	"strconv"

	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model

	Name          string
	PassWord      string
	Phone         string `valid:"matches(^1[3-9]\\d{9}$)"`
	Email         string `valid:"email"`
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
func (table *UserBasic) String() string {
	s := "ID: " + strconv.Itoa(int(table.ID)) + "\n" +
		"Name: " + table.Name + "\n" +
		"PassWord: " + table.PassWord + "\n" +
		"Phone: " + table.Phone + "\n" +
		"Email: " + table.Email + "\n" +
		"Identity: " + table.Identity + "\n" +
		"ClientIp: " + table.ClientIp + "\n" +
		"ClientPort: " + table.ClientPort + "\n" +
		"LoginTime: " + table.LoginTime + "\n" +
		"LoginOutTime: " + table.LoginOutTime + "\n" +
		"HeartbeatTime: " + table.HeartbeatTime + "\n" +
		"IsLogout: " + strconv.FormatBool(table.IsLogout) + "\n" +
		"DeviceInfo: " + table.DeviceInfo + "\n"
	return s
}

// 增
func CreateUser(DB *gorm.DB, user *UserBasic) error {
	// 创建记录
	result := DB.Create(user)
	if result.Error != nil {
		// 如果发生错误，返回错误
		return result.Error
	}

	// 创建成功，返回 nil 错误
	return nil
}
func CreateUserList(DB *gorm.DB, users []*UserBasic) error {
	// 创建记录
	result := DB.Create(users)
	if result.Error != nil {
		// 如果发生错误，返回错误
		return result.Error
	}

	// 创建成功，返回 nil 错误
	return nil
}

// 查
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
func GetUserByID(DB *gorm.DB, id int) (*UserBasic, error) {
	// 初始化一个 UserBasic 对象
	user := &UserBasic{}

	// 执行查询
	result := DB.First(user, id)
	if result.Error != nil {
		// 如果发生错误，返回错误
		return nil, result.Error
	}

	// 查询成功，返回结果和 nil 错误
	return user, nil
}
func GetUsersByName(DB *gorm.DB, name string) ([]*UserBasic, error) {
	// 初始化一个切片来存储查询结果
	var users []*UserBasic

	// 执行查询
	result := DB.Where("name = ?", name).Find(&users)
	if result.Error != nil {
		// 如果发生错误，返回错误
		return nil, result.Error
	}

	// 查询成功，返回结果和 nil 错误
	return users, nil
}
func GetUserByPhone(DB *gorm.DB, phone string) (*UserBasic, error) {
	// 初始化一个 UserBasic 对象
	user := &UserBasic{}

	// 执行查询
	result := DB.Where("phone = ?", phone).First(user)
	if result.Error != nil {
		// 如果发生错误，返回错误
		return nil, result.Error
	}

	// 查询成功，返回结果和 nil 错误
	return user, nil
}

// 改
// 更新用户信息 (根据 ID)
func UpdateUser(DB *gorm.DB, user *UserBasic) error {
	// 更新记录
	result := DB.Save(user)
	if result.Error != nil {
		// 如果发生错误，返回错误
		return result.Error
	}

	// 更新成功，返回 nil 错误
	return nil
}

// 删
func DeleteUserByID(DB *gorm.DB, id int) error {
	// 删除记录
	result := DB.Delete(&UserBasic{}, id)
	if result.Error != nil {
		// 如果发生错误，返回错误
		return result.Error
	}

	// 删除成功，返回 nil 错误
	return nil
}

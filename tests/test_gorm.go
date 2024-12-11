package tests

import (
	"fmt"

	models "github.com/TTaket/gochat/models"
	"github.com/TTaket/gochat/utils"
)

func TestGorm() {
	db := utils.GetDB()

	// 迁移 schema 在数据库中创建或更新相应的表结构。
	db.AutoMigrate(&models.UserBasic{})

	// Create
	user := models.UserBasic{
		Name:       "TestGorm",
		PassWord:   "123456",
		Phone:      "12345678901",
		Email:      "aa@aa.com",
		Identity:   "admin",
		ClientIp:   "192.168.0.1",
		ClientPort: "8080",
		IsLogout:   false,
		DeviceInfo: "Windows 10",
	}
	db.Create(&user)

	// Read
	var userget models.UserBasic
	db.First(&userget, "Name = ?", "TestGorm") // 查找 code 字段值为 D42 的记录
	fmt.Println(userget)
	// Update
	db.Model(&userget).Update("PassWord", "TestGorm_Updated_PassWord")

	// // Update - 更新多个字段
	// db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段
	// db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// // Delete - 删除 product
	// db.Delete(&product, 1)
}

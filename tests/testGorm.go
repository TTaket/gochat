package main

import (
	"fmt"
	"gochat/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	db, err := gorm.Open(mysql.Open("testuser:ABCabc123123...@tcp(127.0.0.1:3306)/gochat"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 迁移 schema 在数据库中创建或更新相应的表结构。
	db.AutoMigrate(&models.UserBasic{})

	// Create
	user := models.UserBasic{
		Name:          "L1212",
		PassWord:      "123456",
		Phone:         "12345678901",
		Email:         "aa@aa.com",
		Identity:      "admin",
		ClientIp:      "192.168.0.1",
		ClientPort:    "8080",
		LoginTime:     "2021-01-01 00:00:00",
		LogoutTime:    "2021-01-01 00:00:00",
		HeartbeatTime: "2021-01-01 00:00:00",
		IsLogout:      false,
		DeviceInfo:    "Windows 10",
	}
	db.Create(&user)

	// Read
	var userget models.UserBasic
	db.First(&userget, "Name = ?", "L1212") // 查找 code 字段值为 D42 的记录
	fmt.Println(userget)
	// Update
	db.Model(&userget).Update("Name", "L1213")

	// // Update - 更新多个字段
	// db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段
	// db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// // Delete - 删除 product
	// db.Delete(&product, 1)
}

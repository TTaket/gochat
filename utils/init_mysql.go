package utils

import (
	"log"
	"os"
	"time"

	"github.com/TTaket/gochat/models"
	viper "github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var mysqlDB *gorm.DB

func InitMysql() error {
	var err error
	err = nil
	newlogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	mysqlDB, err = gorm.Open(mysql.Open(viper.GetString("mysql.dsn")), &gorm.Config{Logger: newlogger})
	if err != nil {
		return err
	}

	// 迁移 schema 在数据库中创建或更新相应的表结构。
	mysqlDB.AutoMigrate(&models.UserBasic{})

	return nil
}

func GetDB() *gorm.DB { // TODO add mutex
	if mysqlDB == nil {
		InitMysql()
	}
	return mysqlDB
}

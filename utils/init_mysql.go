package utils

import (
	"log"
	"os"
	"time"

	viper "github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

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

	DB, err = gorm.Open(mysql.Open(viper.GetString("mysql.dsn")), &gorm.Config{Logger: newlogger})
	if err != nil {
		return err
	}
	return nil
}

func GetDB() *gorm.DB { // TODO add mutex
	if DB == nil {
		InitMysql()
	}
	return DB
}

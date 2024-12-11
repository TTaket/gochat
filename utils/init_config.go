package utils

import (
	"fmt"

	viper "github.com/spf13/viper"
)

func InitConfig() error {
	viper.SetConfigName("app")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	fmt.Println("config mysql:", viper.Get("mysql"))
	return nil
}

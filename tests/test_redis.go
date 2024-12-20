package tests

import (
	"fmt"
	"time"

	"github.com/TTaket/gochat/utils"
	"github.com/go-redis/redis"
)

func TestRedis() {
	var client *redis.Client = utils.GetRedisClient()

	nowtime := time.Now().Format("2006-01-02 15:04:05")
	// Set
	err := client.Set("testkey_time", nowtime, 0).Err()
	if err != nil {
		fmt.Println(err)
	}

	// Get
	val, err := client.Get("testkey_time").Result()
	if err != nil {
		fmt.Println(err)
	}
	if val == nowtime {
		fmt.Println("\033[32mGet testkey_time success\033[0m")
	} else {
		fmt.Println("\033[31mGet testkey_time failed\033[0m")
	}
}

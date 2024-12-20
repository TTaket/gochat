package utils

import (
	redis "github.com/go-redis/redis"
	viper "github.com/spf13/viper"
)

// RedisClient
var RedisClient *redis.Client

// InitRedis
func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.addr"),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.db"),
	})
}

// GetRedisClient
func GetRedisClient() *redis.Client {
	if RedisClient == nil {
		InitRedis()
	}
	return RedisClient
}

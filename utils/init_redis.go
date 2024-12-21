package utils

import (
	redis "github.com/go-redis/redis"
	viper "github.com/spf13/viper"
)

// RedisClient
var redisClient *redis.Client

// InitRedis
func InitRedis() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.addr"),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.db"),
	})
}

// GetRedisClient
func GetRedisClient() *redis.Client {
	if redisClient == nil {
		InitRedis()
	}
	return redisClient
}

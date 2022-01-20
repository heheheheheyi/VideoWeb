package model

import (
	"VideoWeb/config"
	"github.com/go-redis/redis"
	"strconv"
)

var RedisClient *redis.Client

// Redis 在中间件中初始化redis链接
func InitRedis() {
	db, _ := strconv.Atoi(config.RedisDB)
	client := redis.NewClient(&redis.Options{
		Addr:     config.RedisADDR,
		Password: config.RedisPW,
		DB:       db,
	})

	_, err := client.Ping().Result()

	if err != nil {
		panic(err)
	}

	RedisClient = client
}

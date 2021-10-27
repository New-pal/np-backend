package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

type RedisClient struct {
	*redis.Client
}

func NewConnection() *RedisClient {
	client := redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.addr"),
		Password: viper.GetString("redis.pass"),
		DB:       viper.GetInt("redis.db"),
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
	return &RedisClient{client}
}

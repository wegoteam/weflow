package redis

import (
	"context"
	"github.com/wegoteam/weflow/internal/config"

	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client

func Init() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     config.GetConf().Redis.Address,
		Password: config.GetConf().Redis.Password,
	})
	if err := RedisClient.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}
}

package database

import (
	"context"
	"fmt"
	redisconfig "go-api/config/redis_config"

	"github.com/redis/go-redis/v9"
)

func InitRedisDatabase() (*redis.Client, error) {

	config := redisconfig.RedisConfig()

	redisAddr := fmt.Sprintf("%s:%s", config.HOST, config.PORT)

	redis := redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})

	_, err := redis.Ping(context.Background()).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %v", err)
	}

	return redis, nil
}

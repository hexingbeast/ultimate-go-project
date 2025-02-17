package redisdb

import (
	"context"
	"fmt"
	"ultimate-go-project/internal/config"

	"github.com/redis/go-redis/v9"
)

// создается и возвращается новый клиент redis
func NewRedisClient(cfg config.Redis) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr: cfg.Address,
		Password: cfg.Password,
		DB: cfg.DB,
	})

	// проверка соединения
	ctx := context.Background()
	if err := rdb.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("Ошибка подключения в redis: %v", err)
	}

	return rdb, nil
}

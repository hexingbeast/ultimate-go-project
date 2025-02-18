package redisdb

import (
	"context"
	"fmt"
	"ultimate-go-project/internal/config"

	"github.com/redis/go-redis/v9"
)


type RedisStorage struct {
    redisDB *redis.Client
}

// создается и возвращается новый клиент redis
func NewRedisClient(cfg config.Redis) (*RedisStorage, error) {
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

	return &RedisStorage{redisDB: rdb}, nil
}

func (rds *RedisStorage) SaveData(key string, value string) error {
	const op = "storage.redis.SaveData"
	   ctx := context.Background()
	   if err := rds.redisDB.Set(ctx, key, value, 0).Err(); err != nil {
			return fmt.Errorf("%s: %w", op, err)
	   }
	return nil
	//
	//    val, err := rbd.Get(ctx, "testkey").Result()
	//    if err != nil {
	// log.Fatalf("Ошибка при чтении из redis: %v", val)
	//    }
}

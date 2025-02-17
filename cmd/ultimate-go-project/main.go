package main

import (
	"log/slog"
	"os"
	"ultimate-go-project/internal/config"
	"ultimate-go-project/internal/lib/logger"
	redisdb "ultimate-go-project/internal/storage/redis"
)


func main() {
    // создаем конфиг
    cfg := config.MustLoad();
    slog.Info("конфигурация загружена", "config", slog.Any("cfg", cfg))

    // настраиваем уровень логирования в зависимости от энваиромента
    log := logger.Setup(cfg.Env)
    log.Info("ultimate-go-project", slog.String("env", cfg.Env))

    // подключаем redis
    rbd, err := redisdb.NewRedisClient(cfg.Redis)
    if err != nil {
	log.Error("failed to init storage", slog.String("Error", err.Error()))
        os.Exit(1)
    }

    _ = rbd

	//    ctx := context.Background()
	//    // err = rbd.Set(ctx, "testkey", "Hola, redis", 0).Err()
	//    if err := rbd.Set(ctx, "testkey", "Hola, redis", 0).Err(); err != nil {
	// log.Fatalf("Ошибка при записи в redis: %v", err)
	//    }
	//
	//    val, err := rbd.Get(ctx, "testkey").Result()
	//    if err != nil {
	// log.Fatalf("Ошибка при чтении из redis: %v", val)
	//    }
    
}


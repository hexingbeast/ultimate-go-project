package main

import (
	"log/slog"
	"net/http"
	"os"
	"ultimate-go-project/internal/config"
	"ultimate-go-project/internal/lib/logger"
	"ultimate-go-project/internal/router"
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

    // создаем роуты
    router := router.Create(log, rbd)
    log.Info("starting server", slog.String("address", cfg.HTTPServer.Address))

    // создаем сам сервер
    srv := &http.Server{
        Addr: cfg.HTTPServer.Address,
        // router также является handler, получается что это handler
        // внутри с нашими добавленными handler-ами
        Handler: router,
        ReadTimeout: cfg.HTTPServer.Timeout,
        WriteTimeout: cfg.HTTPServer.Timeout,
        IdleTimeout: cfg.HTTPServer.IdleTimeout,
    }

    // вызываем наш сервер, ListenAndServe() это блокирующая функция
    // она не пускает нашу программу дальше
    if err := srv.ListenAndServe(); err != nil {
        log.Error("failed to start server")
    }
    // если сюда программа дошла, то произошла ошибка и сервер остановлен
    log.Error("server stopped")

    _ = rbd
    
}


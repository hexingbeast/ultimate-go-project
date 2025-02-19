package router

import (
	"log/slog"
	"net/http"
	"ultimate-go-project/internal/http-server/handlers/redis/redis_delete"
	"ultimate-go-project/internal/http-server/handlers/redis/redis_get"
	"ultimate-go-project/internal/http-server/handlers/redis/redis_save"
	"ultimate-go-project/internal/lib/logger"
	redisdb "ultimate-go-project/internal/storage/redis"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Create(log *slog.Logger, redisStorage *redisdb.RedisStorage) http.Handler {
    router := chi.NewRouter()

    // middleware for adding requestId in our request(get it from chi dependency)
    router.Use(middleware.RequestID)

    // добавляем handler с кастомным логгером
    router.Use(logger.New(log))

    // если случилась паника внутри handler, не должны останавливать все приложение
    // из-за ошибки в одном handler, восстанавливаем от паники
    router.Use(middleware.Recoverer)
    // надо для того чтобы читать значения из url при подключению к нашему handler
    // привязан к пакету chi
    router.Use(middleware.URLFormat)

    router.Post("/redis", redis_save.New(log, redisStorage))
    router.Get("/redis/{key}", redis_get.GetValue(log, redisStorage))
    router.Delete("/redis/{key}", redis_delete.DeleteValue(log, redisStorage))

    return router 
}

package redis_delete

import (
	"errors"
	"log/slog"
	"net/http"
	"ultimate-go-project/internal/lib/api/response"
	"ultimate-go-project/internal/storage"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

type RedisDataDeletter interface {
    DeleteData(key string) error
}

func DeleteValue(log *slog.Logger, redisDataDeletter RedisDataDeletter) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        const op = "handler.redis.redis_delete.DeleteValue"

        log = log.With(
            slog.String("op", op),
            slog.String("request_id", middleware.GetReqID(r.Context())),
        )

        key := chi.URLParam(r, "key")
		if key == "" {
			log.Info("key is empty")
			render.JSON(w, r, response.Error("invalid request"))
			return
		}

        // удаляем данные из бд
        err := redisDataDeletter.DeleteData(key)
        if errors.Is(err, storage.ErrNotFound) {
            log.Info("key not found", "key", key)
            render.JSON(w, r, response.Error("not found"))
            return
        }
        if err != nil {
            log.Error("failed to delete value", err)
            render.JSON(w, r, response.Error("failed to delete value"))
            return
        }

        log.Info("delete value by key", 
            slog.String("key", key),
        )

        // возвращаем пустой ответ
        w.WriteHeader(http.StatusNoContent)
    }
    
}

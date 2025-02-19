package redis_get

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

type Response struct {
    response.Response
    Key string
    Value string
}

type RedisDataGetter interface {
    GetData(key string) (string, string, error)
}

func GetValue(log *slog.Logger, redisDataGetter RedisDataGetter) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        const op = "handler.redis.redis_get.GetValue"

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

        // получаем данные из бд
        key, value, err := redisDataGetter.GetData(key)
        if errors.Is(err, storage.ErrNotFound) {
            log.Info("key not found", "key", key)
            render.JSON(w, r, response.Error("not found"))
            return
        }
        if err != nil {
            log.Error("failed to get key and value", err)
            render.JSON(w, r, response.Error("failed to get key and value"))
            return
        }

        log.Info("get value by key", 
            slog.String("key", key),
            slog.String("value", value),
        )

        // возвращаем успешный ответ
        render.JSON(w, r, Response{
            Response: response.OK(),
            Key: key,
            Value: value,
        })
    }
    
}

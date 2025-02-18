package redis_save

import (
	"log/slog"
	"net/http"
	"ultimate-go-project/internal/lib/api/response"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
)

type Request struct {
    Key string `json:"key" validate:"required"`
    Value string `json:"value" validate:"required"`
}

type Response struct {
    response.Response
}

type RedisDataSaver interface {
    SaveData(key string, value string) error
}

func New(log *slog.Logger, redisDataSaver RedisDataSaver) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        const op = "handler.redis.redis_save.New"

        log = log.With(
            slog.String("op", op),
            slog.String("request_id", middleware.GetReqID(r.Context())),
        )

        var req Request

        err := render.DecodeJSON(r.Body, &req)
        if err != nil {
            log.Error("failed to decode request body", err)
            render.JSON(w, r, response.Error("failed to decode request"))
            return
        }

        // логгируем распарсенный запрос, для проверки корректности параметров
        log.Info("request body decoded", slog.Any("request", req))

        // валидируем поля в реквесте
        if err := validator.New().Struct(req); err != nil {
            // приводим ошибку к нужному типу
            validateErr := err.(validator.ValidationErrors) 

            // пишем ошибку в лог
            log.Error("invalid request", err)

            // функция resp.ValidationError() формирует человекочитаймую ошибку
            // для ответа на request
            render.JSON(w, r, response.ValidationError(validateErr))
            
            // возвращаем ответ на запрос с ошибкой
            return
        }
        
        // обрабатываем все остальные ошибки
        if err := redisDataSaver.SaveData(req.Key, req.Value); err != nil {
            log.Error("failed to add key and value", err)
            render.JSON(w, r, response.Error("failed to add key and value"))
            return
        }

        // запрос успешно сохранен, запишем это в лог
        log.Info("key and value added", 
            slog.String("key", req.Key),
            slog.String("value", req.Value),
        )

        // возвращаем успешный ответ
        render.JSON(w, r, response.OK())
    }
    
}

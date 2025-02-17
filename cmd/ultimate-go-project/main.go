package main

import (
	"log/slog"
	"ultimate-go-project/internal/config"
)


func main() {
    cfg := config.MustLoad();
    slog.Info("конфигурация загружена", "config", slog.Any("cfg", cfg))
}


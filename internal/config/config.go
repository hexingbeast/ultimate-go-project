package config

import (
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
)
type Config struct {
    Env string `mapstructure:"env"`
    Redis
    HTTPServer `mapstructure:"http_server"`
}

type Redis struct {
    Address string `mapstructure:"address"`
    Password string `mapstructure:"password"`
    DB int `mapstructure:"db"`
}

type HTTPServer struct {
    Address string `mapstructure:"address"`
    Timeout time.Duration `mapstructure:"timeout"`
    IdleTimeout time.Duration `mapstructure:"idle_timeout"`
    User string `mapstructure:"user"`
    Password string `mapstructure:"password"`
}

func MustLoad() *Config {
    configPath := os.Getenv("CONFIG_PATH")
    if configPath == "" {
        log.Fatal("CONFIG_PATH is not set")
    }

    viper.SetConfigFile(configPath)
    viper.SetConfigType("yaml")
    viper.AutomaticEnv()
    
    if err := viper.ReadInConfig(); err != nil {
        log.Fatalf("Ошибка чтения конфигурации: %v", err)
    }

    var cfg Config
    if err := viper.Unmarshal(&cfg); err != nil {
        log.Fatalf("Ошибка при парсинге yaml: %v", err)
    } 

    return &cfg
}


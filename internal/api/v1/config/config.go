package config

import (
	"os"

	"github.com/gen95mis/todo-rest-api/internal/db"
)

// Config ...
type Config struct {
	BindAddr   string
	LogLevel   string
	SessionKey string
	DB         *db.Database
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		BindAddr:   getEnv("BIND_ADDR"),
		LogLevel:   getEnv("LOG_LEVEL"),
		SessionKey: getEnv("SESSION_KEY"),

		DB: &db.Database{
			Host:     getEnv("DB_HOST"),
			Port:     getEnv("DB_PORT"),
			User:     getEnv("DB_USER"),
			Password: getEnv("DB_PASSWORD"),
			Database: getEnv("DB_DATABASE"),
			SSLmode:  getEnv("DB_SSLMODE"),
		},
	}
}

// getEnv получение значения из окружения, если значение отсутствует, то возвращает пустую строку
func getEnv(key string) string {
	value, _ := os.LookupEnv(key)
	return value
}

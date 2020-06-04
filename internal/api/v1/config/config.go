package config

import (
	"github.com/gen95mis/todo-rest-api/internal/db"
	"github.com/gen95mis/todo-rest-api/pkg/env"
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
		BindAddr:   env.GetEnv("BIND_ADDR"),
		LogLevel:   env.GetEnv("LOG_LEVEL"),
		SessionKey: env.GetEnv("SESSION_KEY"),

		DB: &db.Database{
			Host:     env.GetEnv("DB_HOST"),
			Port:     env.GetEnv("DB_PORT"),
			User:     env.GetEnv("DB_USER"),
			Password: env.GetEnv("DB_PASSWORD"),
			Database: env.GetEnv("DB_DATABASE"),
			SSLmode:  env.GetEnv("DB_SSLMODE"),
		},
	}
}

package v1

import "github.com/gen95mis/todo-rest-api/internal/db"

// ConfigAPIServer ...
type ConfigAPIServer struct {
	BindAddr string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
	Database db.Database
}

// NewConfigAPIServer ...
func NewConfigAPIServer() *ConfigAPIServer {
	return &ConfigAPIServer{
		BindAddr: ":8080",
		LogLevel: "debug",
	}
}

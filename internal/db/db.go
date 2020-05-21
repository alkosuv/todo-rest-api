package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // postgres
)

// Database ...
type Database struct {
	Host     string `toml:"host"`
	Port     string `toml:"port"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	Database string `toml:"database"`
	SSLmode  string `toml:"sslmode"`
}

// ConnectDB ...
func (d *Database) ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", d.String())
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func (d *Database) String() string {
	return fmt.Sprintf(
		`host=%s port=%s user=%s password=%s database=%s sslmode=%s`,
		d.Host, d.Port, d.User, d.Password, d.Database, d.SSLmode,
	)
}

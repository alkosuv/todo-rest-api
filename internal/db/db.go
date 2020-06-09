package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // postgres
)

// Database структура для подключения к БД
type Database struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	SSLmode  string
}

// ConnectDB создание подключения к БД
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

// String ...
func (d *Database) String() string {
	return fmt.Sprintf(
		`host=%s port=%s user=%s password=%s database=%s sslmode=%s`,
		d.Host, d.Port, d.User, d.Password, d.Database, d.SSLmode,
	)
}

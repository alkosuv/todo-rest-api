package db

import (
	"database/sql"
	"fmt"
	"strings"
	"testing"

	"github.com/gen95mis/todo-rest-api/pkg/env"
	"github.com/joho/godotenv"
)

// TestConnect подключение к тестовой базе данных
func TestConnect(t *testing.T) (*sql.DB, func(...string)) {
	t.Helper()

	dbConfig := initDB(t)
	db, err := sql.Open("postgres", dbConfig.String())
	if err != nil {
		t.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		t.Fatal(err)
	}

	return db, func(tables ...string) {
		if len(tables) > 0 {
			query := fmt.Sprintf(`truncate %s cascade`, strings.Join(tables, ", "))
			if _, err := db.Exec(query); err != nil {
				t.Fatal(err)
			}
		}

		db.Close()
	}
}

func initDB(t *testing.T) *Database {
	t.Helper()

	if err := godotenv.Load(); err != nil {
		t.Fatal("No .env file found")
	}

	return &Database{
		Host:     env.GetEnv("DB_TEST_HOST"),
		Port:     env.GetEnv("DB_TEST_PORT"),
		User:     env.GetEnv("DB_TEST_USER"),
		Password: env.GetEnv("DB_TEST_PASSWORD"),
		Database: env.GetEnv("DB_TEST_DATABASE"),
		SSLmode:  env.GetEnv("DB_TEST_SSLMODE"),
	}
}

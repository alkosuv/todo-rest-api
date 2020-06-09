package sqlstore

import (
	"database/sql"
	"fmt"

	"github.com/gen95mis/todo-rest-api/internal/api/v1/model"
	"github.com/gen95mis/todo-rest-api/internal/api/v1/store"
)

// UserRepository структура user хранилища
type UserRepository struct {
	db *sql.DB
}

// FindByID поиск пользователя по ID
func (r *UserRepository) FindByID(userID int) (*model.User, error) {
	user := new(model.User)
	err := r.db.QueryRow(
		`SELECT id, login, password, name FROM users WHERE id=$1`,
		userID,
	).Scan(
		&user.ID,
		&user.Login,
		&user.Password,
		&user.Name,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}
	return user, nil
}

// FindByLogin поиск пользователя по login
func (r *UserRepository) FindByLogin(login string) (*model.User, error) {
	user := new(model.User)
	err := r.db.QueryRow(
		`SELECT id, login, password, name FROM users WHERE login=$1`,
		login,
	).Scan(
		&user.ID,
		&user.Login,
		&user.Password,
		&user.Name,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}
	return user, nil
}

// Create создание пользователя
func (r *UserRepository) Create(user *model.User) error {
	return r.db.QueryRow(
		`INSERT INTO users (login, password, name) 
		VALUES ($1, $2, $3) RETURNING id`,
		user.Login, user.Password, user.Name,
	).Scan(&user.ID)
}

// Patch обновление column
func (r *UserRepository) Patch(userID int, column string, value string) error {
	query := fmt.Sprintf(`UPDATE users SET %s=$1 WHERE id=$2`, column)
	if _, err := r.db.Exec(query, value, userID); err != nil {
		return err
	}
	return nil
}

// Exists ...
func (r *UserRepository) Exists(login string, password string) (int, error) {
	var id int
	err := r.db.QueryRow(
		`SELECT id FROM users WHERE login=$1 AND password=$2`,
		login, password,
	).Scan(&id)

	if err != nil {
		if err == sql.ErrNoRows {
			return 0, store.ErrRecordNotFound
		}
		return 0, err
	}

	return id, nil
}

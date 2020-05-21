package sqlstore

import (
	"database/sql"
	"fmt"

	"github.com/gen95mis/todo-rest-api/internal/api/v1/model"
)

// TodoRepository ...
type TodoRepository struct {
	db *sql.DB
}

// GetAll ...
func (r *TodoRepository) GetAll(userID int) ([]*model.Todo, error) {
	rows, err := r.db.Query(
		`SELECT id, title, completed, date_create FROM todos WHERE user_id=$1`,
		userID,
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	todos := []*model.Todo{}
	for rows.Next() {
		todo := new(model.Todo)
		rows.Scan(
			&todo.ID,
			&todo.Title,
			&todo.Completed,
			&todo.DateCreate,
		)
		todos = append(todos, todo)
	}
	return todos, nil
}

// FindCompleted ...
func (r *TodoRepository) FindCompleted(userID int, completed string) ([]*model.Todo, error) {
	rows, err := r.db.Query(
		`SELECT id, title, completed, date_create 
		FROM todos 
		WHERE user_id=$1 AND completed=$2`,
		userID,
		completed,
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	todos := []*model.Todo{}
	for rows.Next() {
		todo := new(model.Todo)
		rows.Scan(
			&todo.ID,
			&todo.Title,
			&todo.Completed,
			&todo.DateCreate,
		)
		todos = append(todos, todo)
	}

	return todos, nil
}

// CountAll ...
func (r *TodoRepository) CountAll(userID int) (int, error) {
	var count int
	err := r.db.QueryRow(
		`SELECT count(*) FROM todos WHERE user_id=$1`,
		userID,
	).Scan(&count)

	if err != nil {
		return 0, err
	}
	return count, nil
}

// CountFindCompleted ...
func (r *TodoRepository) CountCompleted(userID int, completed string) (int, error) {
	var count int
	err := r.db.QueryRow(
		`SELECT count(*) FROM todos WHERE user_id=$1 AND completed=$2`,
		userID, completed,
	).Scan(&count)

	if err != nil {
		return 0, nil
	}

	return count, nil
}

// Create ...
func (r *TodoRepository) Create(todo *model.Todo) error {
	err := r.db.QueryRow(
		`INSERT INTO todos (user_id, title) VALUES ($1, $2) returning id, date_create`,
		todo.UserID,
		todo.Title,
	).Scan(&todo.ID, &todo.DateCreate)

	if err != nil {
		return err
	}

	return nil
}

// Delete ...
func (r *TodoRepository) Delete(userID int, id int) error {
	if _, err := r.db.Exec(
		`DELETE FROM todos WHERE id=$1 AND user_id=$2`,
		id, userID,
	); err != nil {
		return err
	}
	return nil
}

// Patch ...
func (r *TodoRepository) Patch(
	userID int, id int, column string, value string,
) error {
	query := fmt.Sprintf(
		`UPDATE todos SET %s=$1 WHERE id=$2 AND user_id=$3`,
		column,
	)

	if _, err := r.db.Exec(query, value, id, userID); err != nil {
		return err
	}
	return nil
}

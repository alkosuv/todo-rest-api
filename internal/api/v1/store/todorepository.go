package store

import (
	"github.com/gen95mis/todo-rest-api/internal/api/v1/model"
)

// TodoRepository интерфейс описывает функции завпросов к таблице Todo
type TodoRepository interface {
	GetAll(userID int) ([]*model.Todo, error)
	FindByID(userID int, todoID int) (*model.Todo, error)
	FindCompleted(userID int, completed string) ([]*model.Todo, error)
	CountAll(userID int) (int, error)
	CountCompleted(userID int, completed string) (int, error)
	Create(todo *model.Todo) error
	Delete(userID int, todoID int) (bool, error)
	Patch(userID int, todoID int, column string, value string) error
}

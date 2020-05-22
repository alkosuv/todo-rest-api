package store

import (
	"github.com/gen95mis/todo-rest-api/internal/api/v1/model"
)

// UserRepository ...
type UserRepository interface {
	FindByID(id int) (*model.User, error)
	FindByLogin(string) (*model.User, error)
	Create(*model.User) error
	Patch(id int, column string, value string) error
	Exists(login string, password string) (int, error)
}

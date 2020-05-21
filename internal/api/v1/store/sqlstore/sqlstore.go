package sqlstore

import (
	"database/sql"

	"github.com/gen95mis/todo-rest-api/internal/api/v1/store"
)

// SQLStore ...
type SQLStore struct {
	db             *sql.DB
	todoRepository *TodoRepository
	userRepository *UserRepository
}

// NewStore ...
func NewStore(db *sql.DB) *SQLStore {
	return &SQLStore{
		db: db,
	}
}

// Todo ...
func (s *SQLStore) Todo() store.TodoRepository {
	if s.todoRepository != nil {
		return s.todoRepository
	}

	s.todoRepository = &TodoRepository{
		db: s.db,
	}

	return s.todoRepository
}

// User ...
func (s *SQLStore) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		db: s.db,
	}
	return s.userRepository
}

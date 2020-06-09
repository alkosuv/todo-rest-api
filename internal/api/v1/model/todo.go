package model

import "strings"

// Todo ...
type Todo struct {
	ID         int    `json:"id"`
	UserID     int    `json:"user_id,omitempty"`
	Title      string `json:"title"`
	Completed  bool   `json:"completed"`
	DateCreate string `json:"date_create"`
}

// TodoPatchValid валидация patch запросов
func TodoPatchValid(column string, value string) bool {
	t := new(Todo)
	switch strings.ToLower(column) {
	case "title":
		t.Title = value
		return t.IsTitle()
	case "completed":
		switch strings.ToLower(value) {
		case "true", "false":
			return true
		}
	}

	return false
}

// IsTitle валидация заголовка
func (t *Todo) IsTitle() bool {
	if len(t.Title) < 1 || len(t.Title) > 200 {
		return false
	}

	return true
}

package model

import "testing"

// TestUser ...
func TestUser(t *testing.T) *User {
	t.Helper()

	return &User{
		Login:    "user0",
		Password: "12345",
		Name:     "User0 User0",
	}
}

// TestTodo ...
func TestTodo(t *testing.T) *Todo {
	t.Helper()

	return &Todo{
		UserID:    1,
		Title:     "Title",
		Completed: true,
	}
}

package model

import "testing"

// TestUser ...
func TestUser(t *testing.T) *User {
	t.Helper()

	return &User{
		Login:    "user0",
		Password: "12345",
		Name:     "User0",
	}
}

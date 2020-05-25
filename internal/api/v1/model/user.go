package model

import "github.com/gen95mis/todo-rest-api/pkg/validation"

// User ...
type User struct {
	ID       int    `json:"id"`
	Login    string `json:"login"`
	Password string `json:"password,omitempty"`
	Name     string `json:"name"`
}

// IsNil ...
func (u *User) IsNil() bool {
	return *u == User{}
}

// IsUser ...
func (u *User) IsUser() bool {

	return true
}

// IsLogin ...
func (u *User) IsLogin() bool {
	valid, err := validation.IsString(u.Login, 3, 15, `[a-z](\w)`)
	if err != nil {
		return false
	}
	return valid
}

// IsPassword ...
func (u *User) IsPassword() bool {
	valid, err := validation.IsString(u.Password, 5, 15, `\w`)
	if err != nil {
		return false
	}
	return valid
}

// IsName ...
func (u *User) IsName() bool {
	valid, err := validation.IsString(u.Name, 3, 30, `\w`)
	if err != nil {
		return false
	}
	return valid
}

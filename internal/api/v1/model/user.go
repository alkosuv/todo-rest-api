package model

import (
	"strings"

	"github.com/gen95mis/todo-rest-api/pkg/validation"
)

// User ...
type User struct {
	ID       int    `json:"id"`
	Login    string `json:"login"`
	Password string `json:"password,omitempty"`
	Name     string `json:"name"`
}

// UserPatchValid функция для валидации данных для метода patch
func UserPatchValid(column string, value string) bool {
	u := new(User)
	switch strings.ToLower(column) {
	case "password":
		u.Password = value
		return u.IsPassword()
	case "name":
		u.Name = value
		return u.IsName()
	}

	return false
}

// IsNil стуктура User равна nil
func (u *User) IsNil() bool {
	return *u == User{}
}

// IsUser валидаия полей структуры
func (u *User) IsUser() bool {
	return u.IsLogin() && u.IsPassword() && u.IsLogin()
}

// IsLogin валидация поля User.Login
func (u *User) IsLogin() bool {
	valid, err := validation.IsString(u.Login, `^[a-zA-Z][a-z0-9_-]{3,15}$`)
	if err != nil {
		return false
	}
	return valid
}

// IsPassword валидация поля User.Password
func (u *User) IsPassword() bool {
	valid, err := validation.IsString(u.Password, `^[a-zA-Z0-9]{5,15}$`)
	if err != nil {
		return false
	}
	return valid
}

// IsName валидация поля User.Name
func (u *User) IsName() bool {
	valid, err := validation.IsString(u.Name, `^[a-zA-Z0-9 ]{3,30}$`)
	if err != nil {
		return false
	}
	return valid
}

package model_test

import (
	"strings"
	"testing"

	"github.com/gen95mis/todo-rest-api/internal/api/v1/model"
	"github.com/stretchr/testify/assert"
)

func TestUser_ValidLogin(t *testing.T) {
	testCase := []struct {
		name    string
		u       func() *model.User
		isValid bool
	}{
		{
			name: "valid",
			u: func() *model.User {
				return model.TestUser(t)
			},
			isValid: true,
		},
		{
			name: "empty",
			u: func() *model.User {
				return &model.User{}
			},
			isValid: false,
		},
		{
			name: "not valued pattern",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Login = "@@!^@#$"
				return u
			},
			isValid: false,
		},
		{
			name: "fist char number",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Login = "12345"
				return u
			},
			isValid: false,
		},
		{
			name: "lower login",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Login = "ab"
				return u
			},
			isValid: false,
		},
		{
			name: "upper login",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Login = strings.Repeat("a", 20)
				return u
			},
			isValid: false,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			valid := tc.u().IsLogin()
			assert.Equal(t, tc.isValid, valid)
		})
	}
}

func TestUser_ValidPassword(t *testing.T) {
	testCase := []struct {
		name    string
		u       func() *model.User
		isValid bool
	}{
		{
			name: "valid",
			u: func() *model.User {
				return model.TestUser(t)
			},
			isValid: true,
		},
		{
			name: "empty",
			u: func() *model.User {
				return &model.User{}
			},
			isValid: false,
		},
		{
			name: "not valued pattern",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Password = "@@!^@#$"
				return u
			},
			isValid: false,
		},
		{
			name: "lower password",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Password = "abcd"
				return u
			},
			isValid: false,
		},
		{
			name: "upper password",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Password = strings.Repeat("a", 20)
				return u
			},
			isValid: false,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			valid := tc.u().IsPassword()
			assert.Equal(t, tc.isValid, valid)
		})
	}
}

func TestUser_ValidName(t *testing.T) {
	testCase := []struct {
		name    string
		u       func() *model.User
		isValid bool
	}{
		{
			name: "valid",
			u: func() *model.User {
				return model.TestUser(t)
			},
			isValid: true,
		},
		{
			name: "empty",
			u: func() *model.User {
				return &model.User{}
			},
			isValid: false,
		},
		{
			name: "not valued name",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Name = "@@!^@#$"
				return u
			},
			isValid: false,
		},
		{
			name: "lower name",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Name = "ab"
				return u
			},
			isValid: false,
		},
		{
			name: "upper name",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Name = strings.Repeat("a", 35)
				return u
			},
			isValid: false,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			valid := tc.u().IsName()
			assert.Equal(t, tc.isValid, valid)
		})
	}
}

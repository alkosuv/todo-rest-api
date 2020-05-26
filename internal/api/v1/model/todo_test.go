package model_test

import (
	"strings"
	"testing"

	"github.com/gen95mis/todo-rest-api/internal/api/v1/model"
	"github.com/stretchr/testify/assert"
)

func TestTodo_ValidTitle(t *testing.T) {
	testCase := []struct {
		name    string
		t       func() *model.Todo
		isValid bool
	}{
		{
			name: "valid",
			t: func() *model.Todo {
				return model.TestTodo(t)
			},
			isValid: true,
		},
		{
			name: "upper",
			t: func() *model.Todo {
				t := model.TestTodo(t)
				t.Title = strings.Repeat("a", 201)
				return t
			},
			isValid: false,
		},
		{
			name: "lower",
			t: func() *model.Todo {
				t := model.TestTodo(t)
				t.Title = ""
				return t
			},
			isValid: false,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			valid := tc.t().IsTitle()
			assert.Equal(t, tc.isValid, valid)
		})
	}
}

func TestTodo_TodoPatchValid(t *testing.T) {
	testCase := []struct {
		name    string
		column  string
		value   string
		isValid bool
	}{
		{
			name:    "title valid",
			column:  "title",
			value:   "value value",
			isValid: true,
		},
		{
			name:    "title lower",
			column:  "title",
			value:   "",
			isValid: false,
		},
		{
			name:    "title upper",
			column:  "title",
			value:   strings.Repeat("a", 201),
			isValid: false,
		},

		{
			name:    "completed valid true",
			column:  "completed",
			value:   "true",
			isValid: true,
		},
		{
			name:    "completed valid false",
			column:  "completed",
			value:   "false",
			isValid: true,
		},
		{
			name:    "completed any word",
			column:  "completed",
			value:   "word",
			isValid: false,
		},

		{
			name:    "invalid column",
			column:  "column",
			value:   "value",
			isValid: false,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			valid := model.TodoPatchValid(tc.column, tc.value)
			assert.Equal(t, tc.isValid, valid)
		})
	}
}

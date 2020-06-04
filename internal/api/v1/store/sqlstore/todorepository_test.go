package sqlstore_test

import (
	"strconv"
	"testing"

	"github.com/gen95mis/todo-rest-api/internal/api/v1/model"
	"github.com/gen95mis/todo-rest-api/internal/api/v1/store"
	"github.com/gen95mis/todo-rest-api/internal/api/v1/store/sqlstore"
	"github.com/gen95mis/todo-rest-api/internal/db"

	"github.com/stretchr/testify/assert"
)

func TestTodoRepository_Create(t *testing.T) {
	db, teardown := db.TestConnect(t)
	defer teardown("todos")

	todo := model.TestTodo(t)

	s := sqlstore.NewStore(db)
	err := s.Todo().Create(todo)

	assert.NoError(t, err)
	assert.NotNil(t, todo)
}

func TestTodoRepository_Patch(t *testing.T) {
	db, teardown := db.TestConnect(t)
	defer teardown("todos")

	s := sqlstore.NewStore(db)
	todo := model.TestTodo(t)
	value := true

	//case-1
	if err := s.Todo().Create(todo); err != nil {
		t.Fatal(err)
	}
	err := s.Todo().
		Patch(todo.UserID, todo.ID, "completed", strconv.FormatBool(value))
	assert.NoError(t, err)

	todos, err := s.Todo().GetAll(todo.UserID)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, value, todos[0].Completed)

}

func TestTodoRepository_GetAll(t *testing.T) {
	db, teardown := db.TestConnect(t)
	defer teardown("todos")

	todo := model.TestTodo(t)
	s := sqlstore.NewStore(db)

	// case-1
	todos, err := s.Todo().GetAll(todo.UserID)
	assert.Equal(t, 0, len(todos))
	assert.NoError(t, err)

	// case-2
	if err := s.Todo().Create(todo); err != nil {
		t.Fatal(err)
	}
	todos, err = s.Todo().GetAll(todo.UserID)
	assert.NotNil(t, todos)
	assert.NoError(t, err)
}

func TestTodoRepository_FindByID(t *testing.T) {
	db, teardown := db.TestConnect(t)
	defer teardown("todos")

	todo := model.TestTodo(t)
	s := sqlstore.NewStore(db)

	// case-1
	newTodo, err := s.Todo().FindByID(todo.UserID, 1)
	assert.Nil(t, newTodo)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	// case-2
	if err := s.Todo().Create(todo); err != nil {
		t.Fatal(err)
	}
	newTodo, err = s.Todo().FindByID(todo.UserID, todo.ID)
	assert.NotNil(t, newTodo)
	assert.NoError(t, err)

}

func TestTodoRepository_FindCompleted(t *testing.T) {
	db, teardown := db.TestConnect(t)
	defer teardown("todos")

	todo := model.TestTodo(t)
	s := sqlstore.NewStore(db)

	//case-1
	if err := s.Todo().Create(todo); err != nil {
		t.Fatal(err)
	}
	todos, err := s.Todo().FindCompleted(todo.UserID, "true")
	assert.Equal(t, 0, len(todos))
	assert.NoError(t, err)

	// case-2
	todos, err = s.Todo().FindCompleted(todo.UserID, "false")
	assert.NotNil(t, todos)
	assert.NoError(t, err)

	//case-3
	value := true
	err = s.Todo().
		Patch(todo.UserID, todo.ID, "completed", strconv.FormatBool(value))
	if err != nil {
		t.Fatal(err)
	}
	todos, err = s.Todo().FindCompleted(todo.UserID, "true")
	assert.NotNil(t, todos)
	assert.NoError(t, err)
}

func TestTodoRepository_CountAll(t *testing.T) {
	db, teardown := db.TestConnect(t)
	defer teardown("todos")

	todo := model.TestTodo(t)
	s := sqlstore.NewStore(db)

	// case-1
	count, err := s.Todo().CountAll(todo.UserID)
	assert.Equal(t, 0, count)
	assert.NoError(t, err)

	// case-2
	if err := s.Todo().Create(todo); err != nil {
		t.Fatal(err)
	}
	count, err = s.Todo().CountAll(todo.UserID)
	assert.Equal(t, 1, count)
	assert.NoError(t, err)
}

func TestTodoRepository_CountCompleted(t *testing.T) {
	db, teardown := db.TestConnect(t)
	defer teardown("todos")

	todo := model.TestTodo(t)
	s := sqlstore.NewStore(db)
	value := false

	// case-1
	count, err := s.Todo().CountCompleted(todo.UserID, strconv.FormatBool(value))
	assert.Equal(t, 0, count)
	assert.NoError(t, err)

	// case-2
	if err := s.Todo().Create(todo); err != nil {
		t.Fatal(err)
	}
	count, err = s.Todo().CountCompleted(todo.UserID, strconv.FormatBool(value))
	assert.Equal(t, 1, count)
	assert.NoError(t, err)

	// case-3
	value = true
	err = s.Todo().
		Patch(todo.UserID, todo.ID, "completed", strconv.FormatBool(value))
	if err != nil {
		t.Fatal(err)
	}
	count, err = s.Todo().CountCompleted(todo.UserID, strconv.FormatBool(value))
	assert.Equal(t, 1, count)
	assert.NoError(t, err)
}

func TestTodoRepository_Delete(t *testing.T) {
	db, teardown := db.TestConnect(t)
	defer teardown("todos")

	todo := model.TestTodo(t)
	s := sqlstore.NewStore(db)

	//case-1
	if err := s.Todo().Create(todo); err != nil {
		t.Fatal(err)
	}
	ok, err := s.Todo().Delete(todo.UserID, todo.ID)
	assert.True(t, ok)
	assert.NoError(t, err)
}

package sqlstore_test

import (
	"testing"

	"github.com/gen95mis/todo-rest-api/internal/api/v1/model"
	"github.com/gen95mis/todo-rest-api/internal/api/v1/store"
	"github.com/gen95mis/todo-rest-api/internal/api/v1/store/sqlstore"
	"github.com/gen95mis/todo-rest-api/internal/db"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	db, teardown := db.TestConnect(t)
	defer teardown("users")

	user := model.TestUser(t)
	s := sqlstore.NewStore(db)

	// case-1
	err := s.User().Create(user)
	assert.NotNil(t, user.ID)
	assert.NoError(t, err)
}

func TestUserRepository_FindByID(t *testing.T) {
	db, teardown := db.TestConnect(t)
	defer teardown("users")

	user := model.TestUser(t)
	s := sqlstore.NewStore(db)

	// case-1
	u, err := s.User().FindByID(1)
	assert.Nil(t, u)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	// case-2
	if err := s.User().Create(user); err != nil {
		t.Fatal(err)
	}
	u, err = s.User().FindByID(user.ID)
	assert.NotNil(t, u)
	assert.NoError(t, err)
}

func TestUserRepository_FindByLogin(t *testing.T) {
	db, teardown := db.TestConnect(t)
	defer teardown("users")

	user := model.TestUser(t)
	s := sqlstore.NewStore(db)

	// case-1
	u, err := s.User().FindByLogin(user.Login)
	assert.Nil(t, u)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	// case-2
	if err := s.User().Create(user); err != nil {
		t.Fatal(err)
	}
	u, err = s.User().FindByLogin(user.Login)
	assert.NotNil(t, u)
	assert.NoError(t, err)
}

func TestUserRepository_Patch(t *testing.T) {
	db, teardown := db.TestConnect(t)
	defer teardown("users")

	user := model.TestUser(t)
	s := sqlstore.NewStore(db)

	if err := s.User().Create(user); err != nil {
		t.Fatal(err)
	}

	// case-1
	password := "password"
	err := s.User().Patch(user.ID, password, password)
	assert.NoError(t, err)

	u, err := s.User().FindByID(user.ID)
	assert.Equal(t, password, u.Password)
	assert.NoError(t, err)

	// case-2
	name := "name"
	err = s.User().Patch(user.ID, name, name)
	assert.NoError(t, err)

	u, err = s.User().FindByID(user.ID)
	assert.Equal(t, name, u.Name)
	assert.NoError(t, err)
}

func TestUserRepository_Exists(t *testing.T) {
	db, teardown := db.TestConnect(t)
	defer teardown("users")

	user := model.TestUser(t)
	s := sqlstore.NewStore(db)

	// case-1
	id, err := s.User().Exists(user.Login, user.Password)
	assert.Equal(t, 0, id)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	// case-2
	if err := s.User().Create(user); err != nil {
		t.Fatal(err)
	}
	id, err = s.User().Exists(user.Login, user.Password)
	assert.Equal(t, user.ID, id)
	assert.NoError(t, err)
}

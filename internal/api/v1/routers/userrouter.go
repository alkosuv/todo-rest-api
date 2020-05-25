package routers

import (
	"encoding/json"
	"net/http"

<<<<<<< 61c8efc38422bf61a314b8c1c1be0141007a8ca5
	"github.com/gen95mis/todo-rest-api/internal/api/v1/model"
=======
<<<<<<< 5a94ea24bd541a184c7b9038b9b1e7bb19c70e7e
=======
<<<<<<< HEAD
	"github.com/gen95mis/todo-rest-api/internal/api/v1/middleware"
=======
>>>>>>> parent of 5a94ea2... upd middleware
	"github.com/gen95mis/todo-rest-api/internal/api/v1/model"
>>>>>>> upd middleware
>>>>>>> upd middleware
	"github.com/gen95mis/todo-rest-api/internal/api/v1/response"
	"github.com/gen95mis/todo-rest-api/internal/api/v1/store"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// UserRouter ...
type UserRouter struct {
	router *mux.Router
	logger *logrus.Logger
	store  store.Store
}

// NewUserRouter ...
func NewUserRouter(router *mux.Router, logger *logrus.Logger, store store.Store) *UserRouter {
	return &UserRouter{
		router: router,
		logger: logger,
		store:  store,
	}
}

func (ur *UserRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ur.router.ServeHTTP(w, r)
}

// ConfigureRouter ...
func (ur *UserRouter) ConfigureRouter() {
	ur.router.HandleFunc("/users", ur.handlerUserGet()).Methods(http.MethodGet)
	ur.router.HandleFunc("/users", ur.handlerUserPost()).Methods(http.MethodPost)
	ur.router.HandleFunc("/users", ur.handlerUserPatch()).Methods(http.MethodPatch)
}

func (ur *UserRouter) handlerUserGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := 16
		user, err := ur.store.User().FindByID(userID)
		if err != nil {
			response.Error(w, r, http.StatusBadRequest, err)
			return
		}

		response.Response(w, r, http.StatusOK, user)
	}
}

func (ur *UserRouter) handlerUserPost() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := new(model.User)
		json.NewDecoder(r.Body).Decode(user)

		u, _ := ur.store.User().FindByLogin(user.Login)
		if u != nil {
			response.Error(w, r, http.StatusBadRequest, response.ErrLoginUnavailable)
			return
		}

		if err := ur.store.User().Create(user); err != nil {
			response.Error(w, r, http.StatusBadRequest, err)
			return
		}

		response.Response(w, r, http.StatusCreated, user)
	}
}

func (ur *UserRouter) handlerUserPatch() http.HandlerFunc {
	type request struct {
		Column string `json:"column"`
		Value  string `json:"value"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := new(request)
		json.NewDecoder(r.Body).Decode(req)

		userID := 16
		if err := ur.store.User().Patch(userID, req.Column, req.Value); err != nil {
			response.Error(w, r, http.StatusBadRequest, err)
			return
		}

		response.Response(w, r, http.StatusOK, nil)
	}
}

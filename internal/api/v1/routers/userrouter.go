package routers

import (
	"encoding/json"
	"net/http"

	"github.com/gen95mis/todo-rest-api/internal/api/v1/middleware"
	"github.com/gen95mis/todo-rest-api/internal/api/v1/model"
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
	ur.router.HandleFunc("/users", ur.handlerUserPatch()).Methods(http.MethodPatch)
}

func (ur *UserRouter) handlerUserGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(middleware.CtxKeyUser).(*model.User)
		response.Response(w, http.StatusOK, user)
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

		user := r.Context().Value(middleware.CtxKeyUser).(*model.User)

		if err := ur.store.User().Patch(user.ID, req.Column, req.Value); err != nil {
			response.Error(w, http.StatusBadRequest, err)
			return
		}

		response.Response(w, http.StatusOK, nil)
	}
}

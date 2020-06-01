package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gen95mis/todo-rest-api/internal/api/v1/ctxkey"
	"github.com/gen95mis/todo-rest-api/internal/api/v1/log"

	"github.com/gen95mis/todo-rest-api/internal/api/v1/model"
	"github.com/gen95mis/todo-rest-api/internal/api/v1/response"
	"github.com/gen95mis/todo-rest-api/internal/api/v1/store"
	"github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

// TodoRouter ....
type TodoRouter struct {
	router *mux.Router
	logger *logrus.Logger
	store  store.Store
}

// NewTodoRouter ...
func NewTodoRouter(router *mux.Router, logger *logrus.Logger, store store.Store) *TodoRouter {
	return &TodoRouter{
		router: router,
		logger: logger,
		store:  store,
	}
}

func (tr *TodoRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tr.router.ServeHTTP(w, r)
}

// ConfigureRouter ...
func (tr *TodoRouter) ConfigureRouter() {
	tr.router.HandleFunc("/todos", tr.handlerTodosGet()).Methods(http.MethodGet)
	tr.router.HandleFunc("/todos/count", tr.handlerTodosCount()).Methods(http.MethodGet)
	tr.router.HandleFunc("/todos/find", tr.handlerTodosGetCompleted()).Methods(http.MethodGet)
	tr.router.HandleFunc("/todos/find/count", tr.handlerTodosGetCountCompleted()).Methods(http.MethodGet)
	tr.router.HandleFunc("/todos", tr.handlerTodoPost()).Methods(http.MethodPost)
	tr.router.HandleFunc("/todos/{id:[0-9]+}", tr.handlerTodoDelete()).Methods(http.MethodDelete)
	tr.router.HandleFunc("/todos/{id:[0-9]+}", tr.handlerTodoPatch()).Methods(http.MethodPatch)
}

func (tr *TodoRouter) handlerTodosGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(ctxkey.CtxKeyUser).(*model.User)

		todos, err := tr.store.Todo().GetAll(user.ID)
		if err != nil {
			response.Error(w, http.StatusInternalServerError, err)
			log.Error(tr.logger, r, http.StatusInternalServerError, err)
			return
		}

		response.Response(w, http.StatusOK, todos)
		log.Info(tr.logger, r, http.StatusOK, todos)
	}
}

func (tr *TodoRouter) handlerTodosGetCompleted() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		completed := r.URL.Query()["completed"][0]

		user := r.Context().Value(ctxkey.CtxKeyUser).(*model.User)

		todos, err := tr.store.Todo().FindCompleted(user.ID, completed)
		if err != nil {
			response.Error(w, http.StatusInternalServerError, err)
			log.Error(tr.logger, r, http.StatusInternalServerError, err)
			return
		}

		response.Response(w, http.StatusOK, todos)
		log.Info(tr.logger, r, http.StatusOK, todos)
	}
}

func (tr *TodoRouter) handlerTodosCount() http.HandlerFunc {
	type resp struct {
		Count int `json:"count"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(ctxkey.CtxKeyUser).(*model.User)

		count, err := tr.store.Todo().CountAll(user.ID)
		if err != nil {
			response.Error(w, http.StatusInternalServerError, err)
			log.Error(tr.logger, r, http.StatusInternalServerError, err)
			return
		}

		response.Response(w, http.StatusOK, &resp{Count: count})
		log.Info(tr.logger, r, http.StatusOK, &resp{Count: count})
	}
}

func (tr *TodoRouter) handlerTodosGetCountCompleted() http.HandlerFunc {
	type resp struct {
		Count int `json:"count"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		completed := r.URL.Query()["completed"][0]

		user := r.Context().Value(ctxkey.CtxKeyUser).(*model.User)

		count, err := tr.store.Todo().CountCompleted(user.ID, completed)
		if err != nil {
			response.Error(w, http.StatusInternalServerError, err)
			log.Error(tr.logger, r, http.StatusInternalServerError, err)
			return
		}

		response.Response(w, http.StatusOK, &resp{Count: count})
		log.Info(tr.logger, r, http.StatusOK, &resp{Count: count})
	}
}

func (tr *TodoRouter) handlerTodoPost() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		todo := new(model.Todo)
		json.NewDecoder(r.Body).Decode(todo)

		user := r.Context().Value(ctxkey.CtxKeyUser).(*model.User)
		todo.UserID = user.ID

		if !todo.IsTitle() {
			response.Error(w, http.StatusBadRequest, response.ErrIncorrectData)
			log.Error(tr.logger, r, http.StatusBadRequest, response.ErrIncorrectData)
			return
		}

		if err := tr.store.Todo().Create(todo); err != nil {
			response.Error(w, http.StatusUnprocessableEntity, err)
			log.Error(tr.logger, r, http.StatusUnprocessableEntity, err)
			return
		}
		response.Response(w, http.StatusCreated, todo)
		log.Info(tr.logger, r, http.StatusOK, todo)
	}
}

func (tr *TodoRouter) handlerTodoDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(ctxkey.CtxKeyUser).(*model.User)

		id, _ := strconv.Atoi(mux.Vars(r)["id"])
		if err := tr.store.Todo().Delete(user.ID, id); err != nil {
			response.Error(w, http.StatusBadRequest, err)
			log.Error(tr.logger, r, http.StatusBadRequest, err)
			return
		}

		response.Response(w, http.StatusOK, nil)
		log.Info(tr.logger, r, http.StatusOK, nil)
	}
}

func (tr *TodoRouter) handlerTodoPatch() http.HandlerFunc {
	type request struct {
		Column string `json:"column"`
		Value  string `json:"value"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(ctxkey.CtxKeyUser).(*model.User)

		id, _ := strconv.Atoi(mux.Vars(r)["id"])
		req := new(request)
		json.NewDecoder(r.Body).Decode(req)

		if !model.TodoPatchValid(req.Column, req.Value) {
			response.Error(w, http.StatusBadRequest, response.ErrIncorrectData)
			log.Error(tr.logger, r, http.StatusBadRequest, response.ErrIncorrectData)
			return
		}

		if err := tr.store.Todo().Patch(user.ID, id, req.Column, req.Value); err != nil {
			response.Error(w, http.StatusBadRequest, nil)
			log.Error(tr.logger, r, http.StatusBadRequest, err)
			return
		}

		response.Response(w, http.StatusOK, nil)
		log.Info(tr.logger, r, http.StatusOK, nil)
	}
}

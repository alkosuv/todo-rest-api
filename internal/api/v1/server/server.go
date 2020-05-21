package server

import (
	"encoding/json"
	"net/http"

	"github.com/gen95mis/todo-rest-api/internal/api/v1/store"
	"github.com/gorilla/mux"

	"github.com/sirupsen/logrus"
)

// Server ....
type Server struct {
	router *mux.Router
	logger *logrus.Logger
	store  store.Store
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

// NewServer ...
func NewServer(router *mux.Router, logger *logrus.Logger, store store.Store) *Server {
	return &Server{
		router: router,
		logger: logger,
		store:  store,
	}
}

// ConfigureRouter ...
func (s *Server) ConfigureRouter() {
	s.router.HandleFunc("/todos", s.handlerTodosDet()).Methods(http.MethodGet)
	s.router.HandleFunc("/todos/count", s.handlerTodosCount()).Methods(http.MethodGet)
	s.router.HandleFunc("/todos/find", s.handlerTodosGetCompleted()).Methods(http.MethodGet)
	s.router.HandleFunc("/todos/find/count", s.handlerTodosGetCountCompleted()).Methods(http.MethodGet)
	s.router.HandleFunc("/todos", s.handlerTodoPost()).Methods(http.MethodPost)
	s.router.HandleFunc("/todos/{id:[0-9]+}", s.handlerTodoDelete()).Methods(http.MethodDelete)
	s.router.HandleFunc("/todos/{id:[0-9]+}", s.handlerTodoPatch()).Methods(http.MethodPatch)

	s.router.HandleFunc("/user", s.handlerUserGet()).Methods(http.MethodGet)
	s.router.HandleFunc("/user", s.handlerUserPost()).Methods(http.MethodPost)
	s.router.HandleFunc("/user", s.handlerUserPatch()).Methods(http.MethodPatch)
}

/*
response
	...
*/
func (s *Server) response(
	w http.ResponseWriter,
	r *http.Request,
	statusCode int,
	data interface{},
) {
	w.Header().Set("Content-Type", "applications/json")
	w.WriteHeader(statusCode)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

/*
responseError
	...
*/
func (s *Server) responseError(
	w http.ResponseWriter,
	r *http.Request,
	statusCode int,
	err error,
) {
	s.response(w, r, statusCode, map[string]string{"error": err.Error()})
}

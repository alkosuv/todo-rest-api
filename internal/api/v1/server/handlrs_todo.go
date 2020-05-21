package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/gen95mis/todo-rest-api/internal/api/v1/model"
)

func (s *Server) handlerGetTodos() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: UserID инициализировать из Context
		userID := 1
		todos, err := s.store.Todo().GetAll(userID)
		if err != nil {
			s.responseError(w, r, http.StatusInternalServerError, err)
			return
		}

		s.response(w, r, http.StatusOK, todos)
	}
}

func (s *Server) handlerGetTodosSort() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		completed := r.URL.Query()["completed"][0]

		// TODO: UserID инициализировать из Context
		userID := 1
		todos, err := s.store.Todo().FindCompleted(userID, completed)
		if err != nil {
			s.responseError(w, r, http.StatusInternalServerError, err)
			return
		}

		s.response(w, r, http.StatusOK, todos)
	}
}

func (s *Server) handlerTodosCount() http.HandlerFunc {
	type response struct {
		Count int `json:"count"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: UserID инициализировать из Context
		userID := 1
		count, err := s.store.Todo().CountAll(userID)
		if err != nil {
			s.responseError(w, r, http.StatusInternalServerError, err)
		}
		resp := &response{Count: count}
		s.response(w, r, http.StatusOK, resp)
	}
}

func (s *Server) handlerTodosSortCount() http.HandlerFunc {
	type response struct {
		Count int `json:"count"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		completed := r.URL.Query()["completed"][0]

		// TODO: UserID инициализировать из Context
		userID := 1
		count, err := s.store.Todo().CountCompleted(userID, completed)
		if err != nil {
			s.responseError(w, r, http.StatusInternalServerError, err)
		}
		resp := &response{Count: count}
		s.response(w, r, http.StatusOK, resp)
	}
}

func (s *Server) handlerPostTodo() http.HandlerFunc {
	type response struct {
		Title string `json:"title"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		resp := new(response)
		json.NewDecoder(r.Body).Decode(resp)

		// TODO: UserID инициализировать из Context
		todo := &model.Todo{
			Title:  resp.Title,
			UserID: 1,
		}

		if err := s.store.Todo().Create(todo); err != nil {
			s.logger.Error(err)
			s.responseError(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		s.response(w, r, http.StatusCreated, todo)
	}
}

func (s *Server) handlerDeleteTodo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(mux.Vars(r)["id"])
		userID := 1
		if err := s.store.Todo().Delete(userID, id); err != nil {
			s.responseError(w, r, http.StatusBadRequest, err)
			return
		}
		s.response(w, r, http.StatusOK, nil)
	}
}

func (s *Server) handlerPatchTodo() http.HandlerFunc {
	type patch struct {
		Column string `json:"column"`
		Value  string `json:"value"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(mux.Vars(r)["id"])
		// TODO: UserID инициализировать из Context
		userID := 1

		p := &patch{}
		json.NewDecoder(r.Body).Decode(p)

		if err := s.store.Todo().Patch(userID, id, p.Column, p.Value); err != nil {
			s.responseError(w, r, http.StatusBadRequest, nil)
			return
		}

		s.response(w, r, http.StatusOK, nil)
	}
}

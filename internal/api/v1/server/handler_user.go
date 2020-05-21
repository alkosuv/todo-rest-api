package server

import (
	"encoding/json"
	"net/http"

	"github.com/gen95mis/todo-rest-api/internal/api/v1/model"
)

func (s *Server) handlerUserGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := 16
		user, err := s.store.User().FindByID(userID)
		if err != nil {
			s.responseError(w, r, http.StatusBadRequest, err)
			return
		}

		s.response(w, r, http.StatusOK, user)
	}
}

func (s *Server) handlerUserPost() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := new(model.User)
		json.NewDecoder(r.Body).Decode(user)

		u, _ := s.store.User().FindByLogin(user.Login)
		if u != nil {
			s.responseError(w, r, http.StatusBadRequest, ErrLoginUnavailable)
			return
		}

		if err := s.store.User().Create(user); err != nil {
			s.responseError(w, r, http.StatusBadRequest, err)
			return
		}

		s.response(w, r, http.StatusCreated, user)
	}
}

func (s *Server) handlerUserPatch() http.HandlerFunc {
	type response struct {
		Column string `json:"column"`
		Value  string `json:"value"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		resp := new(response)
		json.NewDecoder(r.Body).Decode(resp)

		userID := 16
		if err := s.store.User().Patch(userID, resp.Column, resp.Value); err != nil {
			s.response(w, r, http.StatusBadRequest, err)
			return
		}

		s.response(w, r, http.StatusOK, nil)
	}
}

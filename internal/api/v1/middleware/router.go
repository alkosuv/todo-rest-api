package middleware

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gen95mis/todo-rest-api/internal/api/v1/response"

	"github.com/gen95mis/todo-rest-api/internal/api/v1/store"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/sirupsen/logrus"
)

// Middleware ...
type Middleware struct {
	router  *mux.Router
	logger  *logrus.Logger
	store   store.Store
	session sessions.Store
}

const (
	sessionName = "session"
)

// NewMiddleware  ...
func NewMiddleware(router *mux.Router, logger *logrus.Logger, store store.Store, sessionKey string) *Middleware {
	return &Middleware{
		router:  router,
		logger:  logger,
		store:   store,
		session: sessions.NewCookieStore([]byte(sessionKey)),
	}
}

// ConfigureMiddleware ...
func (m *Middleware) ConfigureMiddleware() {
	m.router.HandleFunc("/sessions", m.handlerSessionCreate()).Methods(http.MethodPost)
}

func (m *Middleware) handlerSessionCreate() http.HandlerFunc {
	type request struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := new(request)
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			response.Error(w, r, http.StatusBadRequest, err)
			return
		}

		// Проверка на существания пользователя
		id, err := m.store.User().Exists(req.Login, req.Password)
		if err != nil {
			response.Error(w, r, http.StatusBadRequest, response.ErrIncorrectLoginOrPassword)
			return
		}

		sessions, err := m.session.Get(r, sessionName)
		if err != nil {
			response.Error(w, r, http.StatusInternalServerError, err)
			return
		}

		sessions.Values["user_id"] = id
		if err := m.session.Save(r, w, sessions); err != nil {
			response.Error(w, r, http.StatusInternalServerError, err)
			return
		}

		response.Response(w, r, http.StatusOK, nil)
	}
}

// AuthenticateUser ...
func (m *Middleware) AuthenticateUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := m.session.Get(r, sessionName)
		if err != nil {
			response.Error(w, r, http.StatusUnauthorized, response.ErrNotAuthenticated)
			return
		}

		id, ok := session.Values["user_id"]
		if !ok {
			response.Error(w, r, http.StatusUnauthorized, response.ErrNotAuthenticated)
		}

		user, err := m.store.User().FindByID(id.(int))
		if err != nil {
			response.Error(w, r, http.StatusUnauthorized, response.ErrNotAuthenticated)
		}

		next.ServeHTTP(w, r.WithContext(context.WithValue(
			r.Context(),
			CtxKeyUser,
			user,
		)))

	})
}

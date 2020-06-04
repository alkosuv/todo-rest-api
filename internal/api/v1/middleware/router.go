package middleware

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gen95mis/todo-rest-api/internal/api/v1/ctxkey"
	"github.com/gen95mis/todo-rest-api/internal/api/v1/log"
	"github.com/gen95mis/todo-rest-api/internal/api/v1/model"
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

// NewMiddleware создание новой Middleware структуры
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
	m.router.HandleFunc("/sessions", m.handlerSessionDelete()).Methods(http.MethodDelete)
	m.router.HandleFunc("/users", m.handlerUserRegister()).Methods(http.MethodPost)
}

// handlerSessionCreate создание сессии
func (m *Middleware) handlerSessionCreate() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		u := new(model.User)
		if err := json.NewDecoder(r.Body).Decode(u); err != nil {
			response.Error(w, http.StatusInternalServerError, nil)
			log.Error(m.logger, r, http.StatusInternalServerError, nil)

			return
		}

		// валидация полученных данных
		if !u.IsLogin() && !u.IsPassword() {
			response.Error(w, http.StatusBadRequest, response.ErrIncorrectLoginOrPassword)
			log.Error(m.logger, r, http.StatusBadRequest, response.ErrIncorrectLoginOrPassword)
			return
		}

		// проверка на существания пользователя
		id, err := m.store.User().Exists(u.Login, u.Password)
		if err != nil {
			response.Error(w, http.StatusBadRequest, response.ErrIncorrectLoginOrPassword)
			log.Error(m.logger, r, http.StatusBadRequest, response.ErrIncorrectLoginOrPassword)
			return
		}

		// создание сессии
		sessions, err := m.session.Get(r, sessionName)
		if err != nil {
			response.Error(w, http.StatusInternalServerError, err)
			log.Error(m.logger, r, http.StatusInternalServerError, err)
			return
		}

		// регистрация пользователя в сессии
		sessions.Values["user_id"] = id
		if err := m.session.Save(r, w, sessions); err != nil {
			response.Error(w, http.StatusInternalServerError, err)
			log.Error(m.logger, r, http.StatusInternalServerError, err)
			return
		}

		response.Response(w, http.StatusOK, nil)
		log.Info(m.logger, r, http.StatusOK, nil)
	}
}

// handlerSessionDelete удаление сессии
func (m *Middleware) handlerSessionDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// получение сессии
		session, err := m.session.Get(r, sessionName)
		if err != nil {
			response.Error(w, http.StatusUnauthorized, response.ErrNotAuthenticated)
			log.Error(m.logger, r, http.StatusUnauthorized, response.ErrNotAuthenticated)
			return
		}
		// удаление сессии
		session.Options.MaxAge = -1
		if err := sessions.Save(r, w); err != nil {
			response.Error(w, http.StatusInternalServerError, err)
			log.Error(m.logger, r, http.StatusUnauthorized, err)
			return
		}
		response.Response(w, http.StatusOK, nil)
		log.Info(m.logger, r, http.StatusOK, nil)
	}
}

// handlerUserRegister регистрация пользователя в системе
func (m *Middleware) handlerUserRegister() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := new(model.User)
		if err := json.NewDecoder(r.Body).Decode(user); err != nil {
			response.Error(w, http.StatusInternalServerError, nil)
			log.Error(m.logger, r, http.StatusInternalServerError, nil)
			return
		}

		// валидация данных
		if !user.IsUser() {
			response.Error(w, http.StatusBadRequest, response.ErrIncorrectData)
			log.Error(m.logger, r, http.StatusBadRequest, response.ErrIncorrectData)
			return
		}

		// проверка на уникальность логина
		if u, _ := m.store.User().FindByLogin(user.Login); u != nil {
			response.Error(w, http.StatusBadRequest, response.ErrLoginUnavailable)
			log.Error(m.logger, r, http.StatusBadRequest, response.ErrLoginUnavailable)
			return
		}

		//  добавление пользователя
		if err := m.store.User().Create(user); err != nil {
			response.Error(w, http.StatusBadRequest, err)
			log.Error(m.logger, r, http.StatusBadRequest, err)
			return
		}

		user.Sanitize()
		response.Response(w, http.StatusCreated, user)
		log.Info(m.logger, r, http.StatusOK, user)
	}
}

// AuthenticateUser авторизация пользователя в системе
func (m *Middleware) AuthenticateUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// получение сессии
		session, err := m.session.Get(r, sessionName)
		if err != nil {
			response.Error(w, http.StatusUnauthorized, response.ErrNotAuthenticated)
			log.Error(m.logger, r, http.StatusUnauthorized, response.ErrNotAuthenticated)
			return
		}

		// получение userID из сессии
		id, ok := session.Values["user_id"]
		if !ok {
			response.Error(w, http.StatusUnauthorized, response.ErrNotAuthenticated)
			log.Error(m.logger, r, http.StatusUnauthorized, response.ErrNotAuthenticated)
			return
		}

		// поиск пользовател в БД по id
		user, err := m.store.User().FindByID(id.(int))
		if err != nil {
			response.Error(w, http.StatusUnauthorized, response.ErrNotAuthenticated)
			log.Error(m.logger, r, http.StatusUnauthorized, response.ErrNotAuthenticated)
			return
		}
		user.Sanitize()

		// запись пользователя в context
		req := r.WithContext(context.WithValue(
			r.Context(),
			ctxkey.CtxKeyUser,
			user,
		))

		next.ServeHTTP(w, req)
	})
}

// UserIsEmpty проверяет наличие авторезованного пользователя в context
func (m *Middleware) UserIsEmpty(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(ctxkey.CtxKeyUser).(*model.User)

		if user.IsNil() {
			response.Response(w, http.StatusInternalServerError, response.ErrSessions)
			log.Error(m.logger, r, http.StatusInternalServerError, response.ErrSessions)
			return
		}

		next.ServeHTTP(w, r)
	})
}

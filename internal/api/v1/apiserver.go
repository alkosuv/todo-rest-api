package v1

import (
	"net/http"
	"os"

	"github.com/gen95mis/todo-rest-api/internal/api/v1/routers"
	"github.com/gen95mis/todo-rest-api/internal/api/v1/store"
	"github.com/gen95mis/todo-rest-api/internal/db"

	"github.com/gen95mis/todo-rest-api/internal/api/v1/middleware"

	"github.com/gen95mis/todo-rest-api/internal/api/v1/store/sqlstore"
	"github.com/sirupsen/logrus"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// APIServer структура api сервера
type APIServer struct {
	BindAddr   string
	LogLevel   string
	SessionKey string
	Database   db.Database
}

// NewAPIServer создание APIServer
func NewAPIServer(bindAddr string, logLevel string, sessionKey string, database *db.Database) *APIServer {
	return &APIServer{
		BindAddr:   bindAddr,
		LogLevel:   logLevel,
		SessionKey: sessionKey,
		Database:   *database,
	}
}

// Start запустить APIServer
func (s *APIServer) Start() error {

	logger := s.initLogger()

	db, err := s.Database.ConnectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	store := sqlstore.NewStore(db)
	router := s.initRouter(logger, store)

	return http.ListenAndServe(s.BindAddr, router)
}

// initRouter инициализация маршруторизатора
func (s *APIServer) initRouter(logger *logrus.Logger, store store.Store) *mux.Router {
	router := mux.NewRouter()
	router.Use(handlers.CORS(handlers.AllowedOrigins([]string{"*"})))

	api := router.PathPrefix("/api/v1").Subrouter()
	api.Use(middleware.SetRequestID)

	m := middleware.NewMiddleware(api, logger, store, s.SessionKey)
	m.ConfigureMiddleware()

	private := api.PathPrefix("/private").Subrouter()
	private.Use(m.AuthenticateUser)
	private.Use(m.UserIsEmpty)

	ur := routers.NewUserRouter(private, logger, store)
	ur.ConfigureRouter()

	tr := routers.NewTodoRouter(private, logger, store)
	tr.ConfigureRouter()

	return router
}

// initLogger инициализация логера
func (s *APIServer) initLogger() *logrus.Logger {
	logger := logrus.New()

	file, err := os.OpenFile("logger.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logger.Out = os.Stdout
	}
	logger.Out = file

	lvl, _ := logrus.ParseLevel(s.LogLevel)
	logger.SetLevel(lvl)
	return logger
}

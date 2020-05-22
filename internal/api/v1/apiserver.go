package v1

import (
	"net/http"

	"github.com/gen95mis/todo-rest-api/internal/api/v1/routers"
	"github.com/gen95mis/todo-rest-api/internal/api/v1/store"
	"github.com/gen95mis/todo-rest-api/internal/db"

	"github.com/gen95mis/todo-rest-api/internal/api/v1/middleware"

	"github.com/gen95mis/todo-rest-api/internal/api/v1/store/sqlstore"
	"github.com/sirupsen/logrus"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// APIServer ...
type APIServer struct {
	BindAddr   string `toml:"bind_addr"`
	LogLevel   string `toml:"log_level"`
	SessionKey string `toml:"session_key"`
	Database   db.Database
}

// NewAPIServer ...
func NewAPIServer() *APIServer {
	return &APIServer{
		BindAddr: ":8080",
		LogLevel: "debug",
	}
}

// Start ...
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

func (s *APIServer) initRouter(logger *logrus.Logger, store store.Store) *mux.Router {
	router := mux.NewRouter()
	router.Use(handlers.CORS(handlers.AllowedOrigins([]string{"*"})))

	api := router.PathPrefix("/api/v1").Subrouter()

	m := middleware.NewMiddleware(api, logger, store, s.SessionKey)
	m.ConfigureMiddleware()

	private := api.PathPrefix("/private").Subrouter()
	private.Use(m.AuthenticateUser)

	ur := routers.NewUserRouter(private, logger, store)
	ur.ConfigureRouter()

	tr := routers.NewTodoRouter(private, logger, store)
	tr.ConfigureRouter()

	return router
}

func (s *APIServer) initLogger() *logrus.Logger {
	logger := logrus.New()
	lvl, _ := logrus.ParseLevel(s.LogLevel)
	logger.SetLevel(lvl)
	return logger
}

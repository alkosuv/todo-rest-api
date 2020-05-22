package v1

import (
	"net/http"

	"github.com/gen95mis/todo-rest-api/internal/api/v1/routers"
	"github.com/gen95mis/todo-rest-api/internal/api/v1/store"

	"github.com/gen95mis/todo-rest-api/internal/api/v1/middleware"

	"github.com/gen95mis/todo-rest-api/internal/api/v1/store/sqlstore"
	"github.com/sirupsen/logrus"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// InitAPIServer ...
func InitAPIServer(config *ConfigAPIServer) error {
	logger := initLogger(config.LogLevel)

	db, err := config.Database.ConnectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	store := sqlstore.NewStore(db)
	router := initRouter(logger, store)

	return http.ListenAndServe(config.BindAddr, router)
}

func initRouter(logger *logrus.Logger, store store.Store) *mux.Router {
	router := mux.NewRouter()
	router.Use(handlers.CORS(handlers.AllowedOrigins([]string{"*"})))

	api := router.PathPrefix("/api/v1").Subrouter()

	m := middleware.NewMiddleware(api, logger, store)
	m.ConfigureMiddleware()

	ur := routers.NewUserRouter(api, logger, store)
	ur.ConfigureRouter()

	tr := routers.NewTodoRouter(api, logger, store)
	tr.ConfigureRouter()

	return router
}

func initLogger(logLevel string) *logrus.Logger {
	logger := logrus.New()
	lvl, _ := logrus.ParseLevel(logLevel)
	logger.SetLevel(lvl)
	return logger
}

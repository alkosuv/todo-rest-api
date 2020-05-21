package v1

import (
	"net/http"

	"github.com/gen95mis/todo-rest-api/internal/api/v1/store/sqlstore"
	"github.com/sirupsen/logrus"

	"github.com/gen95mis/todo-rest-api/internal/api/v1/server"
	"github.com/gorilla/mux"
)

// InitAPIServer ...
func InitAPIServer(config *ConfigAPIServer) error {
	logger := logrus.New()
	lvl, _ := logrus.ParseLevel(config.LogLevel)
	logger.SetLevel(lvl)

	db, err := config.Database.ConnectDB()
	if err != nil {
		logger.Fatal(err.Error())
		return err
	}
	defer db.Close()

	router := mux.NewRouter()
	api := router.PathPrefix("/api/v1").Subrouter()

	store := sqlstore.NewStore(db)

	s := server.NewServer(api, logger, store)
	s.ConfigureRouter()

	return http.ListenAndServe(config.BindAddr, s)
}

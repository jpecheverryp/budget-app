package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/jpecheverryp/budget-app/scs/libsqlstore"
	"github.com/jpecheverryp/budget-app/service"
)

type application struct {
	logger         *slog.Logger
	config         config
	accountService *service.AccountService
	userService    *service.UserService
	sessionManager *scs.SessionManager
}

func main() {
	cfg := getConfig()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	db, err := connectToDB(cfg)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	defer db.Close()

	sessionManager := scs.New()
	sessionManager.Store = libsqlstore.New(db)
	sessionManager.Lifetime = 12 * time.Hour

	app := &application{
		logger:         logger,
		config:         cfg,
		accountService: &service.AccountService{DB: db},
		userService:    &service.UserService{DB: db},
		sessionManager: sessionManager,
	}

	logger.Info("starting server", "port", app.config.port)
	err = http.ListenAndServe(fmt.Sprintf(":%d", app.config.port), app.routes())
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}

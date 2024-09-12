package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/sqlite3store"
	"github.com/alexedwards/scs/v2"
	"github.com/jpecheverryp/budget-app/service"
)

type application struct {
	logger         *slog.Logger
	config         config
	accountService *service.AccountService
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
	sessionManager.Store = sqlite3store.New(db)
	sessionManager.Lifetime = 12 * time.Hour

	app := &application{
		logger:         logger,
		config:         cfg,
		accountService: &service.AccountService{DB: db},
		sessionManager: sessionManager,
	}

	logger.Info("starting server", "port", app.config.port)
	err = http.ListenAndServe(fmt.Sprintf(":%d", app.config.port), app.routes())
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}

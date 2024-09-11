package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/jpecheverryp/budget-app/service"
)

type application struct {
	logger         *slog.Logger
	config         config
	accountService *service.AccountService
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

	app := &application{
		logger:         logger,
		config:         cfg,
		accountService: &service.AccountService{DB: db},
	}

	logger.Info("starting server", "port", app.config.port)
	err = http.ListenAndServe(fmt.Sprintf(":%d", app.config.port), app.routes())
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}

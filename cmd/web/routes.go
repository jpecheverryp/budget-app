package main

import (
	"net/http"

	"github.com/jpecheverryp/budget-app/view"
)

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("GET /static/", http.FileServerFS(view.Files))
	mux.HandleFunc("GET /", app.getIndex)

	mux.HandleFunc("GET /dashboard", app.getDashboard)
	mux.HandleFunc("GET /dashboard/new-account", app.getNewAccount)
	mux.HandleFunc("POST /dashboard/accounts", app.postNewAccount)

	mux.HandleFunc("GET /login", app.getLogin)
	mux.HandleFunc("GET /register", app.getRegister)

	return mux
}

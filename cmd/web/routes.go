package main

import (
	"net/http"

	"github.com/jpecheverryp/budget-app/view"
)

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("GET /static/", http.FileServerFS(view.Files))
	mux.HandleFunc("/", app.getIndex)

	mux.HandleFunc("/dashboard", app.getDashboard)
	mux.HandleFunc("/dashboard/new-account", app.getNewAccount)
	mux.HandleFunc("POST /dashboard/accounts", app.postNewAccount)

	mux.HandleFunc("/login", app.getLogin)
	mux.HandleFunc("/register", app.getRegister)

	return mux
}

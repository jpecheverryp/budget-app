package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", app.getIndex)

	mux.HandleFunc("/dashboard", app.getDashboard)
	mux.HandleFunc("/dashboard/new-account", app.getNewAccount)
	mux.HandleFunc("POST /dashboard/accounts", app.postNewAccount)

	mux.HandleFunc("/login", app.getLogin)
	mux.HandleFunc("/register", app.getRegister)

	return mux
}

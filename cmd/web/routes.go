package main

import (
	"net/http"

	"github.com/jpecheverryp/budget-app/view"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("GET /static/", http.FileServerFS(view.Files))
	mux.HandleFunc("GET /", app.getIndex)

	dynamic := alice.New(app.sessionManager.LoadAndSave)

	mux.Handle("GET /dashboard", dynamic.ThenFunc(app.getDashboard))
	mux.Handle("GET /dashboard/new-account", dynamic.ThenFunc(app.getNewAccount))
	mux.Handle("POST /dashboard/accounts", dynamic.ThenFunc(app.postNewAccount))
	mux.Handle("GET /dashboard/accounts/{id}", dynamic.ThenFunc(app.getAccountInfo))

	mux.Handle("GET /login", dynamic.ThenFunc(app.getLogin))
	mux.Handle("GET /register", dynamic.ThenFunc(app.getRegister))
	mux.Handle("POST /auth/register", dynamic.ThenFunc(app.postRegister))
	mux.Handle("POST /auth/login", dynamic.ThenFunc(app.postLogin))

	standard := alice.New(app.recoverPanic, app.logRequest, commonHeaders)
	return standard.Then(mux)
}

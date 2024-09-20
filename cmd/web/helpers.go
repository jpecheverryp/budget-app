package main

import (
	"net/http"
	"runtime/debug"

	"github.com/a-h/templ"
)

func (app *application) render(w http.ResponseWriter, r *http.Request, component templ.Component) {
	err := component.Render(r.Context(), w)
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) serverError(w http.ResponseWriter, r *http.Request, err error) {
	var (
		method = r.Method
		uri    = r.URL.RequestURI()
		trace  = string(debug.Stack())
	)

	app.logger.Error(err.Error(), "method", method, "uri", uri, "trace", trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *application) isAuthenticated(r *http.Request) bool {
	return app.sessionManager.Exists(r.Context(), "authenticatedUserID")
}

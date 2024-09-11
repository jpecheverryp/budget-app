package main

import (
	"context"
	"net/http"

	"github.com/jpecheverryp/budget-app/view/dashboard"
	"github.com/jpecheverryp/budget-app/view/home"
	"github.com/jpecheverryp/budget-app/view/login"
	"github.com/jpecheverryp/budget-app/view/register"
)

func (app *application) getIndex(w http.ResponseWriter, r *http.Request) {
	home.Show().Render(context.Background(), w)
}

func (app *application) getDashboard(w http.ResponseWriter, r *http.Request) {
	accounts, err := app.accountService.GetAll()
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	dashboard.Show(accounts).Render(context.Background(), w)
}

func (app *application) getLogin(w http.ResponseWriter, r *http.Request) {
	login.Show().Render(context.Background(), w)
}

func (app *application) getRegister(w http.ResponseWriter, r *http.Request) {
	register.Show().Render(context.Background(), w)
}

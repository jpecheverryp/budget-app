package main

import (
	"context"
	"net/http"
	"strconv"

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
		app.serverError(w, r, err)
		return
	}
	dashboard.MainDash(accounts).Render(context.Background(), w)
}

func (app *application) getLogin(w http.ResponseWriter, r *http.Request) {
	login.Show().Render(context.Background(), w)
}

func (app *application) getRegister(w http.ResponseWriter, r *http.Request) {
	register.Show().Render(context.Background(), w)
}

func (app *application) getNewAccount(w http.ResponseWriter, r *http.Request) {
	accounts, err := app.accountService.GetAll()
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	dashboard.ShowNewAccount(accounts).Render(context.Background(), w)

}

func (app *application) postNewAccount(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	accountName := r.PostForm.Get("new-account")

	account, err := app.accountService.Create(accountName)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	dashboard.ShowAccountInfo(account).Render(context.Background(), w)
}

func (app *application) getAccountInfo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	account, err := app.accountService.Read(id)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	// Check if request was made by HTMX and send partial or full data based on that
	isHX := r.Header.Get("HX-Request")
	if isHX == "true" {
		dashboard.ShowAccountInfo(account).Render(context.Background(), w)
	} else {
		accounts, err := app.accountService.GetAll()
		if err != nil {
			app.serverError(w, r, err)
			return
		}
		dashboard.ShowAccountInfoFull(accounts, account).Render(context.Background(), w)
	}

}

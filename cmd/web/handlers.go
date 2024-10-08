package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/jpecheverryp/budget-app/service"
	"github.com/jpecheverryp/budget-app/view/dashboard"
	"github.com/jpecheverryp/budget-app/view/home"
	"github.com/jpecheverryp/budget-app/view/login"
	"github.com/jpecheverryp/budget-app/view/register"
)

type requestData struct {
	userID int
}

func (app *application) newRequestData(r *http.Request) requestData {
	return requestData{
		userID: app.sessionManager.GetInt(r.Context(), "authenticatedUserID"),
	}
}

func (app *application) getIndex(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, home.Show())
}

func (app *application) getDashboard(w http.ResponseWriter, r *http.Request) {
	data := app.newRequestData(r)

	sidebar, err := app.accountService.GetSidebarDataByUserID(data.userID)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	app.render(w, r, dashboard.MainDash(sidebar))
}

func (app *application) getLogin(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, login.Show())
}

func (app *application) postLogin(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	email := r.PostForm.Get("email")
	unencryptedPassword := r.PostForm.Get("password")

	id, err := app.userService.Authenticate(email, unencryptedPassword)
	if err != nil {
		app.logger.Error(err.Error())
		if errors.Is(err, service.ErrInvalidCredentials) {
			app.render(w, r, login.Show())
		} else {
			app.serverError(w, r, err)
		}
		return
	}

	err = app.sessionManager.RenewToken(r.Context())
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	app.sessionManager.Put(r.Context(), "authenticatedUserID", id)

	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}

func (app *application) getRegister(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, register.Show())
}

func (app *application) postRegister(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	username := r.PostForm.Get("username")
	email := r.PostForm.Get("email")
	unencryptedPassword := r.PostForm.Get("password")

	err = app.userService.Insert(username, email, unencryptedPassword)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func (app *application) postLogout(w http.ResponseWriter, r *http.Request) {
	err := app.sessionManager.RenewToken(r.Context())
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	app.sessionManager.Remove(r.Context(), "authenticatedUserID")

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (app *application) getNewAccount(w http.ResponseWriter, r *http.Request) {
	data := app.newRequestData(r)

	sidebar, err := app.accountService.GetSidebarDataByUserID(data.userID)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	app.render(w, r, dashboard.ShowNewAccount(sidebar))
}

func (app *application) postNewAccount(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	accountName := r.PostForm.Get("new-account")
	currentValue, err := strconv.ParseFloat(r.PostForm.Get("current-value"), 64)
	if err != nil {
		app.serverError(w, r, err)
	}
	currentValueInCents := int(currentValue * 100)

	data := app.newRequestData(r)

	account, err := app.accountService.Create(accountName, currentValueInCents, data.userID)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	sidebar, err := app.accountService.GetSidebarDataByUserID(data.userID)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	app.render(w, r, dashboard.ShowAccountInfoFull(sidebar, account))
}

func (app *application) getAccountInfo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	data := app.newRequestData(r)

	account, err := app.accountService.Read(id, data.userID)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	// Check if request was made by HTMX and send partial or full data based on that
	isHX := r.Header.Get("HX-Request")
	if isHX == "true" {
		app.render(w, r, dashboard.ShowAccountInfo(account))
	} else {

		sidebar, err := app.accountService.GetSidebarDataByUserID(data.userID)
		if err != nil {
			app.serverError(w, r, err)
			return
		}

		app.render(w, r, dashboard.ShowAccountInfoFull(sidebar, account))
	}
}

func (app *application) getNewTransaction(w http.ResponseWriter, r *http.Request) {
	data := app.newRequestData(r)

	sidebar, err := app.accountService.GetSidebarDataByUserID(data.userID)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	app.render(w, r, dashboard.ShowNewTransaction(sidebar))
}

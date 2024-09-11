package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/jpecheverryp/budget-app/view/dashboard"
	"github.com/jpecheverryp/budget-app/view/home"
	"github.com/jpecheverryp/budget-app/view/login"
	"github.com/jpecheverryp/budget-app/view/register"
)

type config struct {
	port int
}

type application struct {
	config config
}

func (app *application) getIndex(w http.ResponseWriter, r *http.Request) {
	home.Show().Render(context.Background(), w)
}

func (app *application) getDashboard(w http.ResponseWriter, r *http.Request) {
	dashboard.Show().Render(context.Background(), w)
}

func (app *application) getLogin(w http.ResponseWriter, r *http.Request) {
	login.Show().Render(context.Background(), w)
}

func (app *application) getRegister(w http.ResponseWriter, r *http.Request) {
	register.Show().Render(context.Background(), w)
}

func main() {
	var cfg config
	flag.IntVar(&cfg.port, "port", 8080, "App Network Port")
	flag.Parse()

	app := &application{
		config: cfg,
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", app.getIndex)
	mux.HandleFunc("/dashboard", app.getDashboard)
	mux.HandleFunc("/login", app.getLogin)
	mux.HandleFunc("/register", app.getRegister)

	log.Print("running server in port :", app.config.port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", app.config.port), mux))
}

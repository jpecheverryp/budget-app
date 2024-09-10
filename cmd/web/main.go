package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"

	view "github.com/jpecheverryp/budget-app/view/layout"
)

type config struct {
	port int
}

type application struct {
	config config
}

func (app *application) getIndex(w http.ResponseWriter, r *http.Request) {
    view.Base().Render(context.Background(), w)
}

func (app *application) getDashboard(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This page will show the app dashboard"))
}

func (app *application) getLogin(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This page will show login page"))
}

func (app *application) getRegister(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This page will show register page"))
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

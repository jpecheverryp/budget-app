package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/jpecheverryp/budget-app/service"
	"github.com/jpecheverryp/budget-app/view/dashboard"
	"github.com/jpecheverryp/budget-app/view/home"
	"github.com/jpecheverryp/budget-app/view/login"
	"github.com/jpecheverryp/budget-app/view/register"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

type config struct {
	port int
	dsn  string
}

type application struct {
	config         config
	accountService *service.AccountService
}

func (app *application) getIndex(w http.ResponseWriter, r *http.Request) {
	home.Show().Render(context.Background(), w)
}

func (app *application) getDashboard(w http.ResponseWriter, r *http.Request) {
	accounts, err := app.accountService.GetAll()
	if err != nil {
		log.Print("could not retrieve accounts: ", err)
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

func connectToDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("libsql", cfg.dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func main() {
	var cfg config
	flag.IntVar(&cfg.port, "port", 8080, "App Network Port")
	flag.StringVar(&cfg.dsn, "dsn", "", "Database DSN")
	flag.Parse()

	db, err := connectToDB(cfg)
	if err != nil {
		log.Fatal("cannot connect to database: ", err)
	}
	defer db.Close()

	app := &application{
		config:         cfg,
		accountService: &service.AccountService{DB: db},
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", app.getIndex)
	mux.HandleFunc("/dashboard", app.getDashboard)
	mux.HandleFunc("/login", app.getLogin)
	mux.HandleFunc("/register", app.getRegister)

	log.Print("running server in port :", app.config.port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", app.config.port), mux))
}

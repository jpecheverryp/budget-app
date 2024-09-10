package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

type config struct {
    port int
}

func getIndex (w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("This will be a page to show what the app does"))
}

func getDashboard (w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("This page will show the app dashboard"))
}

func getLogin (w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("This page will show login page"))
}

func getRegister (w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("This page will show register page"))
}

func main() {
    var cfg config
    flag.IntVar(&cfg.port, "port", 8080, "App Network Port")
    flag.Parse()

    mux := http.NewServeMux()

    mux.HandleFunc("/", getIndex)
    mux.HandleFunc("/dashboard", getDashboard)
    mux.HandleFunc("/login", getLogin)
    mux.HandleFunc("/register", getRegister)

    log.Print("running server in port :", cfg.port)
    log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", cfg.port), mux))
}

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

func main() {
    var cfg config
    flag.IntVar(&cfg.port, "port", 8080, "App Network Port")
    flag.Parse()

    mux := http.NewServeMux()

    log.Print("running server in port :", cfg.port)
    log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", cfg.port), mux))
}

package main

import "flag"

type config struct {
	port int
	dsn  string
}

func getConfig() config {
	var cfg config

	flag.IntVar(&cfg.port, "port", 8080, "App Network Port")
	flag.StringVar(&cfg.dsn, "dsn", "", "Database DSN")

	flag.Parse()

	return cfg
}

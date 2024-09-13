package main

import "flag"

type config struct {
	port        int
	dbUrl       string
	dbAuthToken string
}

func getConfig() config {
	var cfg config

	flag.IntVar(&cfg.port, "port", 8080, "App Network Port")
	flag.StringVar(&cfg.dbUrl, "db_url", "", "Database URL")
	flag.StringVar(&cfg.dbAuthToken, "db_auth_token", "", "Database Auth Token")

	flag.Parse()

	return cfg
}

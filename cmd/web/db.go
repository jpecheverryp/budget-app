package main

import (
	"database/sql"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

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

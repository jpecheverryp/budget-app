package main

import (
	"database/sql"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

func connectToDB(cfg config) (*sql.DB, error) {
	dbDsn := cfg.dbUrl + "?authToken=" + cfg.dbAuthToken
	db, err := sql.Open("libsql", dbDsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

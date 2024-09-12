package service

import "database/sql"

type User struct {
	ID                  int
	Username            string
	Email               string
	UnencryptedPassword string
	HashedPassword      string
}

type UserService struct {
	DB *sql.DB
}

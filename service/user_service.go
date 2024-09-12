package service

import (
	"database/sql"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID                  int
	Username            string
	Email               string
	UnencryptedPassword string
	HashedPassword      string
	CreatedAt           time.Time
	UpdatedAt           time.Time
}

type UserService struct {
	DB *sql.DB
}

func (s UserService) Insert(username, email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	stmt := `INSERT INTO users (username, email, hashed_password) VALUES (?, ?, ?)`

	_, err = s.DB.Exec(stmt, username, email, string(hashedPassword))
	if err != nil {
		return err
	}

	return nil
}

func (s UserService) Authenticate(email, password string) (int, error) {
	return 0, nil
}

func (s UserService) Exists(id int) (bool, error) {
	return false, nil
}

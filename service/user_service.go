package service

import (
	"database/sql"
	"errors"
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
	var id int
	var hashedPassword []byte

	stmt := `SELECT rowid, hashed_password FROM users WHERE email = ?`

	err := s.DB.QueryRow(stmt, email).Scan(&id, &hashedPassword)
	if err != nil {
		return 0, err
	}

	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return 0, ErrInvalidCredentials
		} else {
			return 0, err
		}
	}

	return id, nil
}

func (s UserService) GetUsernameByID(id int) (string, error) {
	var username string
	stmt := `SELECT username FROM users WHERE rowid = ?`
	err := s.DB.QueryRow(stmt, id).Scan(&username)
	if err != nil {
		return "", err
	}
	return username, nil
}

func (s UserService) Exists(id int) (bool, error) {
	return false, nil
}

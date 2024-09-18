package service

import (
	"database/sql"
	"time"
)

type Account struct {
	ID          int
	AccountName string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type AccountService struct {
	DB *sql.DB
}

type SidebarData struct {
	ID       int
	Username string
	Accounts []Account
}

func (s AccountService) GetSidebarDataByUserID(id int) (SidebarData, error) {
	var sidebarData SidebarData
	stmt := `SELECT username FROM users WHERE rowid = ?`
	err := s.DB.QueryRow(stmt, id).Scan(&sidebarData.Username)
	if err != nil {
		return SidebarData{}, err
	}

	accounts, err := s.GetAll()
	if err != nil {
		return SidebarData{}, err
	}
	sidebarData.Accounts = accounts

	return sidebarData, nil
}

func (s AccountService) Read(id int) (Account, error) {
	var a Account
	stmt := `SELECT account_name, created_at, updated_at FROM account WHERE rowid = ?`
	err := s.DB.QueryRow(stmt, id).Scan(&a.AccountName, &a.CreatedAt, &a.UpdatedAt)
	if err != nil {
		return Account{}, nil
	}
	a.ID = id
	return a, nil
}

func (s AccountService) GetAll() ([]Account, error) {
	stmt := `SELECT rowid, account_name, created_at, updated_at FROM account`
	rows, err := s.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var accounts []Account
	for rows.Next() {
		var a Account
		err := rows.Scan(
			&a.ID,
			&a.AccountName,
			&a.CreatedAt,
			&a.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		accounts = append(accounts, a)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return accounts, nil
}

func (s AccountService) Create(accountName string) (Account, error) {
	var a Account
	stmt := `INSERT INTO account (account_name) VALUES (?) RETURNING rowid, created_at, updated_at`
	err := s.DB.QueryRow(stmt, accountName).Scan(&a.ID, &a.CreatedAt, &a.UpdatedAt)
	if err != nil {
		return Account{}, nil
	}
	a.AccountName = accountName
	return a, nil
}

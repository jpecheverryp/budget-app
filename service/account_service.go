package service

import (
	"database/sql"
	"time"
)

type Account struct {
	ID           int
	AccountName  string
	CurrentValue int
	CreatedAt    time.Time
	UpdatedAt    time.Time
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
	sidebarData.ID = id
	stmt := `SELECT username FROM users WHERE id = ?`
	err := s.DB.QueryRow(stmt, id).Scan(&sidebarData.Username)
	if err != nil {
		return SidebarData{}, err
	}

	accounts, err := s.GetAll(sidebarData.ID)
	if err != nil {
		return SidebarData{}, err
	}
	sidebarData.Accounts = accounts

	return sidebarData, nil
}

func (s AccountService) Read(id int, userID int) (Account, error) {
	var a Account
	stmt := `SELECT account_name, current_value, created_at, updated_at FROM account WHERE id = ? AND user_id = ?`
	err := s.DB.QueryRow(stmt, id, userID).Scan(&a.AccountName, &a.CurrentValue, &a.CreatedAt, &a.UpdatedAt)
	if err != nil {
		return Account{}, nil
	}
	a.ID = id
	return a, nil
}

func (s AccountService) GetAll(userID int) ([]Account, error) {
	stmt := `SELECT id, account_name, current_value, created_at, updated_at FROM account WHERE user_id = ?`
	rows, err := s.DB.Query(stmt, userID)
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
			&a.CurrentValue,
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

func (s AccountService) Create(accountName string, currentValue int, userID int) (Account, error) {
	var a Account
	stmt := `INSERT INTO account (account_name, current_value, user_id) VALUES (?, ?, ?) RETURNING id, created_at, updated_at`
	err := s.DB.QueryRow(stmt, accountName, currentValue, userID).Scan(&a.ID, &a.CreatedAt, &a.UpdatedAt)
	if err != nil {
		return Account{}, nil
	}
	a.AccountName = accountName
	a.CurrentValue = currentValue
	return a, nil
}

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

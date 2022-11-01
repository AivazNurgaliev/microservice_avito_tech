package data

import (
	"database/sql"
	"errors"
)

type Account struct {
	AccountId           int64   `json:"account_id"`
	AccountCash         float64 `json:"account_cash"`
	AccountReservedCash float64 `json:"account_reserved_cash"`
}

type AccountModel struct {
	DB *sql.DB
}

// Create account with empty parameters
func (a AccountModel) Create() {
	query := `
		INSERT INTO account (account_cash, account_reserved_cash)
		VALUES (0, 0)`

	a.DB.QueryRow(query).Scan()
}

//Create account with certain id
func (a AccountModel) CreateId(id int64) {
	query := `
		INSERT INTO account (account_id, account_cash, account_reserved_cash)
		VALUES ($1, 0, 0)`

	a.DB.QueryRow(query, id).Scan()
}

func (a AccountModel) GetAccountById(id int64) (*Account, error) {
	if id < 1 {
		return nil, errors.New("incorrect id")
	}

	query := `
			SELECT account_id, account_cash, account_reserved_cash
			FROM account
			WHERE account_id = $1
		`

	var account Account
	err := a.DB.QueryRow(query, id).Scan(
		&account.AccountId,
		&account.AccountCash,
		&account.AccountReservedCash,
	)
	if err != nil {
		return nil, errors.New("Account does not exist")
	}

	return &account, nil
}

func (a AccountModel) UpdateAccountCashById(account *Account) error {
	query := `
		UPDATE account
		SET account_cash = $2
		WHERE account_id = $1;`

	return a.DB.QueryRow(query, account.AccountId, account.AccountCash).Scan()
}

func (a AccountModel) UpdateAccountCashData(account *Account) error {
	query := `
		UPDATE account
		SET account_cash = $2, account_reserved_cash = $3
		WHERE account_id = $1;
`

	return a.DB.QueryRow(query, account.AccountId,
		account.AccountCash, account.AccountReservedCash).Scan()
}

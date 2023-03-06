package domain

import (
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/nikzayn/banking/errs"
	"github.com/nikzayn/banking/logger"
)

type AccountRepositoryDb struct {
	client *sqlx.DB
}

func (d AccountRepositoryDb) Save(a Account) (*Account, *errs.AppError) {
	saveQuery := `insert into accounts (customer_id, opening_date, account_type, amount, status) values (?, ?, ?, ?, ?)`

	result, err := d.client.Exec(saveQuery, a.CustomerId, a.OpeningDate, a.AccountType, a.Amount, a.Status)
	if err != nil {
		logger.Error("Error while creating new account: " + err.Error())
		return nil, errs.UnexpectedError("Unexpected database error")
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while last insert id for new account: " + err.Error())
		return nil, errs.UnexpectedError("Unexpected database error")
	}
	a.AccountId = strconv.FormatInt(id, 10)
	return &a, nil
}

func NewAccountRepositoryDb(dbClient *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{dbClient}
}

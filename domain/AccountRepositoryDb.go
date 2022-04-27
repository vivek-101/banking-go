package domain

import (
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/vivek-101/banking/errs"
	"github.com/vivek-101/banking/logger"
)

type AccountRepositoryDb struct {
	client *sqlx.DB
}

func (d AccountRepositoryDb) Save(a Account) (*Account, *errs.AppError) {
	sqlinsert := "INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) values(?, ?, ?, ?, ?)"
	result, err := d.client.Exec(sqlinsert, a.CustomerId, a.OpeningDate, a.AccountType, a.Amount, a.Status)
	if err != nil {
		logger.Error("Error while creating new account:" + err.Error())
		return nil, errs.NewUnexpextedError("Unexpected error from database")
	}
	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting last insert value:" + err.Error())
		return nil, errs.NewUnexpextedError("Unexpected error from database")
	}
	a.AccountId = strconv.FormatInt(id, 10)
	return &a, nil
}

func NewAccountRepositoryDb(dbClient *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{dbClient}
}

package domain

import (
	"banking-app/errs"
	"log"
	"strconv"

	"github.com/jmoiron/sqlx"
)

type AccountRepositoryDb struct {
	client *sqlx.DB
}

func (ar AccountRepositoryDb) Save(a Account) (*Account, *errs.AppError) {
	sql := `
		INSERT INTO accounts (customer_id, opening_date, account_type, amount, status)
		VALUES (?, ?, ?, ?, ?)
	`
	result, err := ar.client.Exec(sql, a.CustomerId, a.OpeningDate, a.AccountType, a.Amount, a.Status)

	if err != nil {
		log.Println("Error while executing INSERT INTO account statement", err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	id, err := result.LastInsertId()

	if err != nil {
		log.Println("Error while getting last inserted id from database")
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	a.AccoundId = strconv.FormatInt(id, 10)

	return &a, nil
}

func NewAccountRepositoryDb(dbClient *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{dbClient}
}

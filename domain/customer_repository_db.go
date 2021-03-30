package domain

import (
	"banking-app/config"
	"banking-app/errs"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func (cr CustomerRepositoryDb) FindAll() ([]Customer, *errs.AppError) {
	findAllSql := `
		SELECT customer_id, name, city, zipcode, date_of_birth, status 
		FROM customers
	`
	customers := make([]Customer, 0)
	err := cr.client.Select(&customers, findAllSql)

	if err != nil {
		log.Println("Error while querying customer table", err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	return customers, nil
}

func (cr CustomerRepositoryDb) FindAllWithStatus(status string) ([]Customer, *errs.AppError) {
	findAllSql := `
		SELECT customer_id, name, city, zipcode, date_of_birth, status 
		FROM customers 
		WHERE status = ?
	`
	customers := make([]Customer, 0)
	err := cr.client.Select(&customers, findAllSql, status)

	if err != nil {
		log.Println("Error while querying customer table", err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	return customers, nil
}

func (cr CustomerRepositoryDb) FindById(id string) (*Customer, *errs.AppError) {
	customerSql := `
		SELECT customer_id, name, city, zipcode, date_of_birth, status 
		FROM customers 
		WHERE customer_id = ?
	`
	var c Customer
	err := cr.client.Get(&c, customerSql, id)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		} else {
			log.Println("Error while scanning customer", err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}

	return &c, nil
}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
	dbAddr := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		config.DB_USERNAME, config.DB_PASSWORD, config.DB_ADDRESS, config.DB_PORT, config.DB_SCHEMA,
	)

	client, err := sqlx.Open("mysql", dbAddr)
	if err != nil {
		panic(err)
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return CustomerRepositoryDb{client}
}

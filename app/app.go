package app

import (
	"banking-app/config"
	"banking-app/domain"
	"banking-app/service"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func Start() {
	if !config.CheckEnvLoaded() {
		log.Fatal("Environmnet variables are not loaded")
		return
	}

	dbClient := createDbConnection()

	customerRepository := domain.NewCustomerRepositoryDb(dbClient)
	accountRepository := domain.NewAccountRepositoryDb(dbClient)

	customerService := service.NewDefaultCustomerService(customerRepository)
	accountService := service.NewDefaultAccountService(accountRepository)

	ch := CustomerHandlers{customerService}
	ah := AccountHandlers{accountService}

	router := mux.NewRouter()

	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customers/{id:[0-9]+}/account", ah.newAccount).Methods(http.MethodPost)

	log.Printf("Started server on port %s", config.SERVER_PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", config.SERVER_ADDRESS, config.SERVER_PORT), router))
}

func createDbConnection() *sqlx.DB {
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

	return client
}

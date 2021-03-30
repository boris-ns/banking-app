package app

import (
	"banking-app/config"
	"banking-app/domain"
	"banking-app/service"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	if !config.CheckEnvLoaded() {
		log.Fatal("Environmnet variables are not loaded")
		return
	}

	ch := CustomerHandlers{service.NewDefaultCustomerService(domain.NewCustomerRepositoryDb())}

	router := mux.NewRouter()

	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	log.Printf("Started server on port %s", config.SERVER_PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", config.SERVER_ADDRESS, config.SERVER_PORT), router))
}

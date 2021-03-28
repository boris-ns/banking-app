package app

import (
	"banking-app/domain"
	"banking-app/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	ch := CustomerHandlers{service.NewDefaultCustomerService(domain.NewCustomerRepositoryDb())}

	router := mux.NewRouter()

	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	log.Println("Started server on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}

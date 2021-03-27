package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	router := mux.NewRouter()

	router.HandleFunc("/greet", greet).Methods(http.MethodGet)
	router.HandleFunc("/customers", getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)
	router.HandleFunc("/customers/{id:[0-9]+}", getCustomer).Methods(http.MethodGet)

	log.Println("Started server on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}

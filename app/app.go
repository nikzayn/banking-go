package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nikzayn/banking/domain"
	"github.com/nikzayn/banking/service"
)

func Start() {
	router := mux.NewRouter()

	//wiring
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	//routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomerById).Methods(http.MethodGet)

	//server
	log.Println(http.ListenAndServe(":8000", router))
}

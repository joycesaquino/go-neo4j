package main

import (
	"dera-services-api/internal/service"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/service", service.NewController().Insert).Methods(http.MethodPost)
}

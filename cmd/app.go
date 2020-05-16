package main

import (
	"dera-services-api/internal/service"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/service", service.Insert).Methods(http.MethodPost)
}

package main

import (
	"dera-services-api/internal/service"
	"github.com/gorilla/mux"
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"net/http"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/service", service.NewController().Insert).Methods(http.MethodPost)
	router.HandleFunc("/service_delete", service.NewController().Create).Methods(http.MethodDelete)

	//API Interface de Programação de Aplicações
	// Camada HTTP

	//	CRUD
	//Create -> POST
	//Read -> GET
	//Update -> PUT
	//Delete -> DELETE

	//SERVICES

}

type Operation struct {
	Value    float64
	Code     string
	Type     string
	DateBuy  neo4j.Date
	DateSale neo4j.Date
	Taxes    float64
	Profit   float64
}

func Calculate(operation Operation) {
	if operation.DateSale.Day() > operation.DateBuy.Day() {
		operation.Taxes = operation.Value * 2
	}
}

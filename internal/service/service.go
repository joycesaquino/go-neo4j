package service

import (
	"fmt"
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"io/ioutil"
	"net/http"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Service struct {
	Id               string              `json:"id"`
	Description      string              `json:"description"`
	Value            float64             `json:"value"`
	InitialDateTime  neo4j.LocalDateTime `json:"date"`
	FinalDateTime    neo4j.LocalDateTime `json:"date"`
	MinSubscriptions int64               `json:"minSubscriptions"`
	MaxSubscriptions int64               `json:"maxSubscriptions"`
	CreatedAt        neo4j.LocalDateTime `json:"createdAt"`
}

const (
	BadRequest = "Bad Request 404"
)

func Insert(w http.ResponseWriter, r *http.Request) {
	var _ Service
	_, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("BAD REQUEST 400")
	}
}

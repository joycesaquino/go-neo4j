package service

import (
	"encoding/json"
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

func Insert(w http.ResponseWriter, r *http.Request) {
	var s *Service
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("Bad Request Error : %s", err)
	}
	u := &User{
		Name:  r.Header.Get("User"),
		Email: r.Header.Get("Email"),
	}

	err = json.Unmarshal(b, &s)
	if err != nil {
		fmt.Printf("Error unmarshaling body to service interface : %s", err)
		w.WriteHeader(http.StatusUnprocessableEntity)
	}

	err = NewDao().Insert(s, u)
	if err != nil {
		fmt.Printf("Error on create service : %s", err)
		w.WriteHeader(http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusCreated)
}

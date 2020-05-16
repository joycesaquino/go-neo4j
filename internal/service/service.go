package service

import (
	"dera-services-api/internal/dao"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Controller struct {
	dao *dao.Dao
}

func (c Controller) Insert(w http.ResponseWriter, r *http.Request) {
	var s *dao.Service
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("Bad Request Error : %s", err)
	}
	u := &dao.User{
		Name:  r.Header.Get("User"),
		Email: r.Header.Get("Email"),
	}

	err = json.Unmarshal(b, &s)
	if err != nil {
		fmt.Printf("Error unmarshaling body to service interface : %s", err)
		w.WriteHeader(http.StatusUnprocessableEntity)
	}

	err = c.dao.Insert(s, u)
	if err != nil {
		fmt.Printf("Error on create service : %s", err)
		w.WriteHeader(http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusCreated)
}

func NewController() *Controller {
	return &Controller{dao: dao.NewDao()}
}

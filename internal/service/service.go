package service

import (
	"dera-services-api/internal/dao"
	"dera-services-api/internal/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Controller struct {
	dao *dao.Dao
}

func (c Controller) Create(w http.ResponseWriter, r *http.Request) {

}

func (c Controller) Insert(w http.ResponseWriter, r *http.Request) {
	var s *model.Class
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("Bad Request Error : %s", err)
	}
	u := &model.User{
		Name:  r.Header.Get("User"),
		Email: r.Header.Get("Email"),
	}

	err = json.Unmarshal(b, &s)
	if err != nil {
		fmt.Printf("Error unmarshaling body to service interface : %s", err)
		w.WriteHeader(http.StatusUnprocessableEntity)
	}

	err = c.dao.InsertClass(s, u.Name)
	if err != nil {
		fmt.Printf("Error on create service : %s", err)
		w.WriteHeader(http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusCreated)
}

func NewController() *Controller {
	return &Controller{dao: dao.NewDao()}
}

package controller

import (
	"encoding/json"
	"net/http"

	"github.com/jijeshmohan/goseed/model"
)

func ListUser(rw http.ResponseWriter, r *http.Request) {
	users := model.GetAllUsers()
	renderJson(rw, users, nil)
}

func CreateUser(rw http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user model.User
	if err := decoder.Decode(&user); err != nil {
		renderJson(rw, nil, err)
	}
	err := model.CreateNewUser(&user)
	renderJson(rw, user, err)
}

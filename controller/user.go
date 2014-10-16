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
	userJson := r.FormValue("user")
	var user model.User
	json.Unmarshal([]byte(userJson), &user)
	err := model.CreateNewUser(&user)
	renderJson(rw, user, err)
}

package controller

import (
	"encoding/json"
	"net/http"

	"github.com/jijeshmohan/goseed/model"
)

func ListUser(rw http.ResponseWriter, r *http.Request) (interface{}, error) {
	users := model.GetAllUsers()
	return users, nil
}

func CreateUser(rw http.ResponseWriter, r *http.Request) (interface{}, error) {
	userJson := r.FormValue("user")
	var user model.User
	json.Unmarshal([]byte(userJson), &user)
	err := model.CreateNewUser(&user)
	return user, err
}

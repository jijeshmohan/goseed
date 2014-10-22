package controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

func DeleteUser(rw http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	userId, _ := strconv.ParseInt(id, 0, 64)
	user := model.GetUser(userId)
	if user == nil {
		renderJson(rw, nil, errors.New("Unable to find user with id "+id))
		return
	}
	model.DeleteUser(user)
	renderJson(rw, true, nil)
}

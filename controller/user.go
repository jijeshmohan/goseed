package controller

import (
	"net/http"

	"github.com/jijeshmohan/goseed/model"
)

func ListUser(rw http.ResponseWriter, r *http.Request) interface{} {
	users := model.GetAllUsers()
	return users
}

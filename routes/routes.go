package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	c "github.com/jijeshmohan/goseed/controller"
)

type Route struct {
	Path   string
	F      http.HandlerFunc
	Method string
}

var routes = []Route{
	{"/users", c.ListUser, "GET"},
	{"/users", c.CreateUser, "POST"},
}

func InitRoutes(router *mux.Router) {
	for _, r := range routes {
		router.Handle(r.Path, http.HandlerFunc(r.F)).Methods(r.Method)
	}
}

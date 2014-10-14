package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	c "github.com/jijeshmohan/goseed/controller"
	m "github.com/jijeshmohan/goseed/middleware"
)

type HttpHandleFunc func(http.ResponseWriter, *http.Request) interface{}

type Route struct {
	Path   string
	F      HttpHandleFunc
	Method string
}

var routes = []Route{
	{"/users", c.ListUser, "GET"},
}

func InitRoutes(router *mux.Router) {
	for _, r := range routes {
		router.Handle(r.Path, m.JsonHandle(r.F)).Methods(r.Method)
	}
}

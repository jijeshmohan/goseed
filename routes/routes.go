package routes

import (
	"net/http"

	c "github.com/jijeshmohan/goseed/controller"
)

type HttpHandleFunc func(http.ResponseWriter, *http.Request) interface{}

type Route struct {
	Path   string
	F      HttpHandleFunc
	Method string
}

var Routes = []Route{
	{"/users", c.ListUser, "GET"},
}

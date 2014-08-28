package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	m "github.com/jijeshmohan/goseed/middleware"
	"github.com/jijeshmohan/goseed/routes"
)

type HttpServer struct {
	router *mux.Router
	uiDir  string
}

func NewHttpServer(uiDir string) *HttpServer {
	return &HttpServer{router: mux.NewRouter(), uiDir: uiDir}
}

func (h *HttpServer) StartServer(address string) {
	h.initRoutes()
	log.Printf("Starting server at %s", address)

	http.ListenAndServe(address, handlers.LoggingHandler(os.Stdout, h.router))
}

func (h *HttpServer) initRoutes() {
	for _, r := range routes.Routes {
		h.router.Handle(r.Path, m.JsonHandler{r.F}).Methods(r.Method)
	}
	h.router.PathPrefix("/").Handler(http.FileServer(http.Dir(h.uiDir)))
}

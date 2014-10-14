package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

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
	h.setupRoutes()
	log.Printf("Starting server at %s", address)
	if err := http.ListenAndServe(address, handlers.LoggingHandler(os.Stdout, h.router)); err != nil {
		log.Fatalln(err)
	}
}

func (h *HttpServer) setupRoutes() {
	routes.InitRoutes(h.router)
	h.router.PathPrefix("/").Handler(http.FileServer(http.Dir(h.uiDir)))
}

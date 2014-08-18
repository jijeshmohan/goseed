package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
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
	http.ListenAndServe(address, h.router)
}

func (h *HttpServer) initRoutes() {
	h.router.PathPrefix("/").Handler(http.FileServer(http.Dir(h.uiDir)))
}

package main

import (
	"encoding/json"
	"log"
	"net/http"

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

type appHandler struct {
	handler func(http.ResponseWriter, *http.Request) interface{}
}

func marshal(item interface{}, w http.ResponseWriter) {
	bytes, err := json.Marshal(item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write(bytes)
}

func (t appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data := t.handler(w, r)
	marshal(data, w)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	log.Println(r.URL, "-", r.Method, "-", r.RemoteAddr)
}

func (h *HttpServer) StartServer(address string) {
	h.initRoutes()
	log.Printf("Starting server at %s", address)
	http.ListenAndServe(address, h.router)
}

func (h *HttpServer) initRoutes() {
	for _, r := range routes.Routes {
		h.router.Handle(r.Path, appHandler{r.F}).Methods(r.Method)
	}
	h.router.PathPrefix("/").Handler(http.FileServer(http.Dir(h.uiDir)))
}

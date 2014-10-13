package middleware

import (
	"encoding/json"
	"log"
	"net/http"
)

type JsonHandler struct {
	Handler func(http.ResponseWriter, *http.Request) interface{}
}

func (t JsonHandler) marshal(item interface{}, w http.ResponseWriter) {
	bytes, err := json.Marshal(item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write([]byte(")]}',")) // This is to avoid angularjs JSON Vulnerability Protection(https://docs.angularjs.org/api/ng/service/$http)
	w.Write(bytes)
}

func (t JsonHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data := t.Handler(w, r)
	t.marshal(data, w)
	log.Println(r.URL, "-", r.Method, "-", r.RemoteAddr)
}

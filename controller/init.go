package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func renderJson(w http.ResponseWriter, data interface{}, err error) {
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintln(err)))
		return
	}
	marshal(data, w)
}

func marshal(item interface{}, w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(")]}',\n")) // This is to avoid angularjs JSON Vulnerability Protection(https://docs.angularjs.org/api/ng/service/$http)
	json.NewEncoder(w).Encode(item)
}

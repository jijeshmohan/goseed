package middleware

import (
	"encoding/json"
	"net/http"
)

func JsonHandle(handle func(http.ResponseWriter, *http.Request) interface{}) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data := handle(w, r)
		marshal(data, w)
	})
}

func marshal(item interface{}, w http.ResponseWriter) {
	bytes, err := json.Marshal(item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(")]}',\n")) // This is to avoid angularjs JSON Vulnerability Protection(https://docs.angularjs.org/api/ng/service/$http)
	w.Write(bytes)
}

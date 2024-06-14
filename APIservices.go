package main

import (
	"encoding/json"
	"net/http"
)

type APIFunc func(w http.ResponseWriter, r *http.Request) error

// @JavaComparison : It is like the Exception handler

type ApiError struct {
	Error string `json:"error"`
}

// @USAGE  : func is passed by the value to the function
// Make it the Higher Order function
func makeHttpHandler(f APIFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if err := f(writer, request); err != nil {
			err = WriteJson(writer, http.StatusBadRequest, &ApiError{Error: err.Error()})
			if err != nil {
				return
			}
		}
	}
}

func WriteJson(w http.ResponseWriter, status int, message any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(message)
}

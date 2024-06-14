package main

import "net/http"

type AuthenticationController interface {
	handleLogin(w http.ResponseWriter, r http.Request) error
	handleRegister(w http.ResponseWriter, r http.Request) error
}

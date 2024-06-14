package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

type APIServer struct {
	ListeningAddress string
}

func NewAPIServer(listenAddress string) *APIServer {
	return &APIServer{
		ListeningAddress: listenAddress,
	}
}

func (api *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/account", makeHttpHandler(api.handleAccount))

	err := http.ListenAndServe(api.ListeningAddress, nil)
	if err != nil {
		return
	}
}

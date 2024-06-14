package main

import "net/http"

type AccountControllerHandler interface {
	handleAccount(w http.ResponseWriter, r *http.Request) error
	handleCreateAccount(w http.ResponseWriter, r *http.Request) error
	handleGetAccount(w http.ResponseWriter, r *http.Request) error
	handleGetAccountById(w http.ResponseWriter, r *http.Request) error
	handleDeleteAccount(w http.ResponseWriter, r *http.Request) error
	handleUpdateAccount(w http.ResponseWriter, r *http.Request) error
}

func (api *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
	return WriteJson(w, http.StatusOK, "Successful API server")
}
func (api *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (api *APIServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func (api *APIServer) handleGetAccountById(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (api *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (api *APIServer) handleUpdateAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

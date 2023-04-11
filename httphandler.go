package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type HTTPHandler struct {
	// Will can add a database here to have access to your data.
	// For example:
	// database DataRepository
	// and then you create a type `DataRepository` with functions
	// to read/store your data :)
}

func (h *HTTPHandler) Get(rw http.ResponseWriter, request *http.Request) {
	rw.WriteHeader(http.StatusOK)
	jsonOutput := map[string]any{
		"hello": "stranger",
	}
	json.NewEncoder(rw).Encode(jsonOutput)
}

func (h *HTTPHandler) GetWithParams(rw http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	name := vars["name"]

	rw.WriteHeader(http.StatusOK)
	jsonOutput := map[string]any{
		"hello": name,
	}
	json.NewEncoder(rw).Encode(jsonOutput)
}

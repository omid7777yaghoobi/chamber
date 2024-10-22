package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func ConfigureServerHandler() (http.Handler, error) {
	router := mux.NewRouter()

	// Registring different Handlers
	RegisterApiHandler(router)

	return router, nil
}

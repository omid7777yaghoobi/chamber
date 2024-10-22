package cmd

import (
	"github.com/gorilla/mux"
)

func RegisterApiHandler(router *mux.Router) {
	apiRouter := router.NewRoute().PathPrefix("/api/v4").Subrouter()

	apiRouter.Handle("/bucket", bucketHandler)
}

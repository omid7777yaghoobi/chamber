package main

import (
	"github.com/gorilla/mux"
)

const (
	apiVersion      = "v1"
	apiPrefix       = "/api/" + apiVersion
	bucketApiPrefix = apiPrefix + "/buckets"
	objectApiPrefix = apiPrefix + "/objects"
)

func RegisterApiHandler(router *mux.Router) {
	apiRouter := router.NewRoute().PathPrefix(apiPrefix).Subrouter()

	registerBucketApiHandler(apiRouter)
	registerObjectApiRouter(apiRouter)
}

func registerBucketApiHandler(router *mux.Router) {
	bucketApiRouter := router.NewRoute().PathPrefix(bucketApiPrefix).Subrouter()

}

func registerObjectApiRouter(router *mux.Router) {
	objectApiRouter := router.NewRoute().PathPrefix(objectApiPrefix).Subrouter()
}

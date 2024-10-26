package main

import (
	"net/http"

	"chamber.io/internal/bucket"
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
	// registerObjectApiRouter(apiRouter)
}

func registerBucketApiHandler(router *mux.Router) {
	bucketApiRouter := router.NewRoute().PathPrefix(bucketApiPrefix).Subrouter()

	bucketApiRouter.Methods(http.MethodGet).Path("/buckets").HandlerFunc(bucket.ListBucketsHandler())

}

// func registerObjectApiRouter(router *mux.Router) {
// 	objectApiRouter := router.NewRoute().PathPrefix(objectApiPrefix).Subrouter()
// }

func ListBucketsHandler() http.HandlerFunc {
	var listBucketsHandler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		// request parser
		// call the actual function
	}

	return listBucketsHandler
}

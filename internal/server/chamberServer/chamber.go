package chamberServer

import (
	"net/http"

	"chamber.io/src/bucket"
)

type ChamberServer struct {
	config  string
	version string
	routes  []Route
}

type UrlPrefix string

type Route struct {
	routePath    string
	routeHandler http.Handler
}

type ChamberServerRoutes struct {
	AuthRoutes   map[UrlPrefix]http.Handler
	BucketRoutes map[UrlPrefix]http.Handler
	ObjectRoutes map[UrlPrefix]http.Handler
}

func ChamberServerInit() {

	sag := http.ServeMux()
}

func (cs *ChamberServer) withAuthRoutes() {
	bucketApiMux := bucket.NewBucketApi()
	authServer := NewAuthServer()
	authRoutes := Route{"/auth", authServer}
	bucketRoutes := Route{"/buckets", bucketApiMux}
	cs.routes = append(cs.routes)
}

type AuthServer struct {
}

func NewAuthServer() http.Handler {
	khar := AuthServer{}
	return &khar
}

func (as *AuthServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	return
}

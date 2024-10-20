package bucket

import (
	"net/http"
)

type bucketAPI struct {
}

func NewBucketApi() http.Handler {
	ListBuckets := NewListBuckets()
	mux := http.NewServeMux()
	mux.Handle("/buckets", ListBuckets)

	return mux
}

type ListBuckets struct {
}

func NewListBuckets() http.Handler {
	buckets := ListBuckets{}
	return &buckets
}

func (lb *ListBuckets) ServeHTTP(r http.ResponseWriter, w *http.Request) {

}

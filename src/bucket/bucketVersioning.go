package bucket

import (
	"net/http"
)

type BucketVersioning struct {
}

func (bv *BucketVersioning) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	return
}

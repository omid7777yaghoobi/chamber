package api

import (
	"chamber.io/internal/bucket"
)

type BucketListMessageBody struct {
	buckets []bucket.Bucket
}

func NewBucketListMessageBody(buckets []bucket.Bucket) *BucketListMessageBody {
	return &BucketListMessageBody{buckets: buckets}
}

func (blm *BucketListMessageBody) Serialize() ([]byte, error) {

}

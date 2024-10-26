package bucket

import (
	"net/http"

	"chamber.io/internal/api"
	"chamber.io/internal/bucket"
	"github.com/google/uuid"
)

func ListBucketsHandler() http.HandlerFunc {
	var listBucketsHandler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		// request parser
		// call the actual function
		correlationId := uuid.New()
		msgAttrs := api.NewMessageAttributes(correlationId.String())
		// msgBody := bucket.ListBuckets()
		msgBody := api.NewBucketListMessageBody(bucket.ListBuckets()).Serialize()

		res := api.NewApiResponse().
			WithMessageAttributes(*msgAttrs).
			WithMessageBody(msgBody)

		serializedResponse, err := res.Serialize()
		if err != nil {
			w.Write("fucked up")
		}
		w.Write(serializedResponse)
	}

	return listBucketsHandler
}

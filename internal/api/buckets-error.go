package api

import (
	"net/http"
)

// type ApiError struct {
// 	Method int
// 	// Headers []http.Header
// 	Body []byte
// }

// func (ae *ApiError) WithApiErrorMethod(method int) *ApiError {
// 	ae.Method = method
// 	return ae
// }

// func (ae *ApiError) WithApiErrorBody(body apiErrorBody) *ApiError {
// 	errBody, _ := json.Marshal(body)
// 	ae.Body = errBody
// 	return ae
// }

// type apiErrorBody struct {
// 	Status  string
// 	Message string
// }

func BucketNotFound() *ApiError {
	apiError := &ApiError{}
	return apiError.WithApiErrorMethod(http.StatusNotFound).
		WithApiErrorBody(apiErrorBody{
			Status:  "Error",
			Message: "Bucket not Found.",
		})
}

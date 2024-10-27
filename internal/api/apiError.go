package api

import (
	"encoding/json"
	"net/http"
)

type ApiError struct {
	Method int
	// Headers []http.Header
	Body []byte
}

func (ae *ApiError) WithApiErrorMethod(method int) *ApiError {
	ae.Method = method
	return ae
}

func (ae *ApiError) WithApiErrorBody(body apiErrorBody) *ApiError {
	errBody, _ := json.Marshal(body)
	ae.Body = errBody
	return ae
}

type apiErrorBody struct {
	Status  string
	Message string
}

func ChamberInternalServerError() *ApiError {
	apiError := &ApiError{}
	return apiError.WithApiErrorMethod(http.StatusInternalServerError).
		WithApiErrorBody(apiErrorBody{
			Status:  "Error",
			Message: "Internal server error.",
		})
}

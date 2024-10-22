package main

import (
	"fmt"

	chttp "chamber.io/internal/http"
)

func main() {
	serverAddr := ""

	serverHandler, err := ConfigureServerHandler()
	if err != nil {
		fmt.Println("Handler initialization failed.")
		return
	}

	chamberServer := chttp.NewServer(serverAddr).
		WithHandler(serverHandler)

	chamberServer.ListenAndServe()
}

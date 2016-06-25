package main

import (
	"net/http"

	dcgi "github.com/bfirsh/go-dcgi"
	"github.com/docker/engine-api/client"
)

func main() {

	// Create client using environment variables, if present
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	// Setup default catch-all handler
	http.Handle("/", &dcgi.Handler{
		Image:  "root",
		Client: cli,
	})

	// Setup specific asset handler, trailing slash is important
	http.Handle("/assets/", &dcgi.Handler{
		Image:  "assets",
		Client: cli,
	})

	// Start listener
	http.ListenAndServe(":8080", nil)
}

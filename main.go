package main

import (
	"net/http"

	dcgi "github.com/bfirsh/go-dcgi"
	"github.com/docker/engine-api/client"
)

func main() {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	http.Handle("/", &dcgi.Handler{
		Image:  "root",
		Client: cli,
	})
	http.Handle("/assets/", &dcgi.Handler{
		Image:  "assets",
		Client: cli,
	})
	http.ListenAndServe(":8080", nil)
}

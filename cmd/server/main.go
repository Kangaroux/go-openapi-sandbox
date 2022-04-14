package main

import (
	"log"
	"net/http"

	api "github.com/kangaroux/go-openapi-test"
	_ "github.com/kangaroux/go-openapi-test/docs"
)

const (
	port = "8000"
)

func main() {
	r := api.NewAPIRouter()

	log.Printf("Listening on port %s", port)

	if err := http.ListenAndServe("0.0.0.0:"+port, r); err != nil {
		log.Fatal(err)
	}
}

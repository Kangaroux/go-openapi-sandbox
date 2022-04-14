package main

import (
	"log"
	"net/http"
	"os"

	_ "github.com/kangaroux/go-openapi-test/docs"
)

func main() {
	r := NewAPIRouter()
	port := os.Getenv("SERVER_PORT")

	if port == "" {
		port = "8000"
	}

	log.Printf("Listening on port %s", port)

	if err := http.ListenAndServe("0.0.0.0:"+port, r); err != nil {
		log.Fatal(err)
	}
}

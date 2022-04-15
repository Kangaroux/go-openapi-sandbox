package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"

	"github.com/jmoiron/sqlx"
	api "github.com/kangaroux/go-openapi-sandbox"
	_ "github.com/kangaroux/go-openapi-sandbox/docs"
)

const (
	port = "8000"
)

func main() {
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s/%s?sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_DB"),
	)
	db, err := sqlx.Connect("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	apis := api.RouterAPIs{
		User: api.NewUserAPI(db),
	}

	r := api.NewAPIRouter(apis)

	log.Printf("Listening on port %s", port)

	if err := http.ListenAndServe("0.0.0.0:"+port, r); err != nil {
		log.Fatal(err)
	}
}

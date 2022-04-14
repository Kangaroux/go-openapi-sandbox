package api

import (
	"github.com/gorilla/mux"
)

func NewAPIRouter() *mux.Router {
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()

	api.PathPrefix("/users").Handler(NewUserRouter("/api/v1/users"))

	return r
}

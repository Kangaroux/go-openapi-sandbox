package api

import (
	"github.com/gorilla/mux"
)

type RouterAPIs struct {
	User UserAPI
}

func NewAPIRouter(apis RouterAPIs) *mux.Router {
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()

	api.PathPrefix("/users").Handler(apis.User.Router("/api/v1/users"))

	return r
}

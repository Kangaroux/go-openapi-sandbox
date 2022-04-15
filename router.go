package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

type RouterAPIs struct {
	User UserAPI
}

func NewAPIRouter(apis RouterAPIs) *mux.Router {
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()

	api.Path("/schema").HandlerFunc(OpenAPISchemaView)
	api.PathPrefix("/users").Handler(apis.User.Router("/api/v1/users"))

	return r
}

// swagger:route GET /schema getSchema
//
// Returns a yaml of the OpenAPI schema
//
// Produces:
//   - application/yaml
func OpenAPISchemaView(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("content-type", "application/yaml")
	http.ServeFile(w, req, "swagger.yml")
}

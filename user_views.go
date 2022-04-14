package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func NewUserRouter(prefix string) *mux.Router {
	r := mux.NewRouter()

	// List
	r.Path(prefix).Name("create user").Methods("POST").HandlerFunc(CreateUser)
	r.Path(prefix).Name("list users").Methods("GET").HandlerFunc(ListUsers)

	// Detail
	r.Path(prefix + `/{id:\d+}`).Name("get user").Methods("GET").HandlerFunc(GetUser)

	return r
}

// swagger:route GET /users users listUsers
//
// Responses:
//   200: listUserResponse
func ListUsers(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("ListUsers"))
}

// swagger:route POST /users users createUser
//
// Responses:
//   200: userResponse
func CreateUser(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("CreateUser"))
}

// swagger:route GET /users/{id} users getUser
//
// Responses:
//   200: userResponse
func GetUser(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	w.Write([]byte(fmt.Sprintf("GetUser: %s", id)))
}

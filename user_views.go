package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

type UserAPI struct {
	db    *sqlx.DB
	users UserService
}

func NewUserAPI(db *sqlx.DB) UserAPI {
	return UserAPI{
		db:    db,
		users: NewUserService(db),
	}
}

func (api UserAPI) Router(prefix string) *mux.Router {
	r := mux.NewRouter()

	// List
	r.Path(prefix).Name("create user").Methods("POST").HandlerFunc(api.CreateUser)
	r.Path(prefix).Name("list users").Methods("GET").HandlerFunc(api.ListUsers)

	// Detail
	r.Path(prefix + `/{id:\d+}`).Name("get user").Methods("GET").HandlerFunc(api.GetUser)

	return r
}

// swagger:route GET /users users listUsers
//
// Responses:
//   200: listUserResponse
func (api UserAPI) ListUsers(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("ListUsers"))
}

// swagger:route POST /users users createUser
//
// Responses:
//   200: userResponse
func (api UserAPI) CreateUser(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("CreateUser"))
}

// swagger:route GET /users/{id} users getUser
//
// Responses:
//   200: userResponse
func (api UserAPI) GetUser(w http.ResponseWriter, req *http.Request) {
	id, err := strconv.ParseInt(mux.Vars(req)["id"], 10, 64)

	if err != nil {
		log.Fatal(err)
	}

	resp := UserResponse{}
	u, err := api.users.Get(id)

	if err != nil {
		log.Fatal(err)
	} else if u == nil {
		resp.Body.OK = false
		resp.Body.Error = "user does not exist"
		w.WriteHeader(404)
	} else {
		resp.Body.OK = true
		resp.Body.User = u
	}

	serialized, err := json.Marshal(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(serialized)
}

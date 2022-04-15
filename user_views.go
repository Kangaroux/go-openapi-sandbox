package api

import (
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
	if err := req.ParseForm(); err != nil {
		log.Print(err)
		WriteJSON(w, ErrorResponse("unable to parse form: "+err.Error()), 400)
		return
	}

	WriteJSON(w, OKResponse())
}

// swagger:route GET /users/{id} users getUser
//
// Consumes:
//   - application/json
//
// Produces:
//   - application/json
//
// Parameters:
//   + name: id
//     in: path
//     type: integer
//     required: true
//
// Responses:
//   200: userResponse
//   404: baseResponse
//   500: baseResponse
func (api UserAPI) GetUser(w http.ResponseWriter, req *http.Request) {
	id, err := strconv.ParseInt(mux.Vars(req)["id"], 10, 64)

	if err != nil {
		log.Print(err)
		WriteJSON(w, InternalErrorResponse(), 500)
		return
	}

	u, err := api.users.Get(id)

	if err != nil {
		log.Print(err)
		WriteJSON(w, InternalErrorResponse(), 500)
	} else if u == nil {
		WriteJSON(w, ErrorResponse("user does not exist"), 404)
	} else {
		resp := UserResponse{}
		resp.Body.OK = true
		resp.Body.User = u
		WriteJSON(w, resp)
	}
}

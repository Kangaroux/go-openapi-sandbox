package api

import (
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

var regexEmail *regexp.Regexp

func init() {
	regexEmail = regexp.MustCompile(`^.+@.+$`)
}

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
	type createUserForm struct {
		Email    string `json:"email"`
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var err error
	form := createUserForm{}

	if !ParseRequestJSON(w, req, &form) {
		return
	}

	form.Email = strings.TrimSpace(form.Email)
	form.Username = strings.TrimSpace(form.Username)

	fieldErrors := make(map[string]string)

	var emailInUse, usernameInUse bool

	if form.Email == "" {
		fieldErrors["email"] = "cannot be blank"
	} else if !regexEmail.MatchString(form.Email) {
		fieldErrors["email"] = "must be a valid email address"
	} else if emailInUse, err = api.users.Exists("email = $1", form.Email); err != nil {
		log.Print(err)
		WriteJSON(w, NewInternalErrorResponse(), 500)
		return
	} else if emailInUse {
		fieldErrors["email"] = "email is already in use"
	}

	if form.Username == "" {
		fieldErrors["username"] = "cannot be blank"
	} else if len(form.Username) < 3 || len(form.Username) > 16 {
		fieldErrors["username"] = "username must be between 3 and 16 characters"
	} else if usernameInUse, err = api.users.Exists("username = $1", form.Username); err != nil {
		log.Print(err)
		WriteJSON(w, NewInternalErrorResponse(), 500)
		return
	} else if usernameInUse {
		fieldErrors["username"] = "username is already in use"
	}

	if len(fieldErrors) > 0 {
		WriteJSON(w, NewFieldErrorResponse(fieldErrors, "some fields need to be corrected").Body, 400)
	} else {
		WriteJSON(w, form)
	}
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
		WriteJSON(w, NewInternalErrorResponse(), 500)
		return
	}

	u, err := api.users.Get(id)

	if err != nil {
		log.Print(err)
		WriteJSON(w, NewInternalErrorResponse(), 500)
	} else if u == nil {
		WriteJSON(w, NewErrorResponse("user does not exist"), 404)
	} else {
		resp := UserResponse{}
		resp.Body.OK = true
		resp.Body.User = u
		WriteJSON(w, resp)
	}
}

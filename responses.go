package api

// swagger:response baseResponse
type BaseResponse struct {
	// An error message describing what went wrong
	Error string `json:"error,omitempty"`

	// Whether the request was successful or not
	OK bool `json:"ok"`
}

// swagger:response userResponse
type UserResponse struct {
	// in: body
	Body struct {
		// swagger:allOf
		BaseResponse

		User *User `json:"user,omitempty"`
	}
}

// swagger:response listUserResponse
type ListUserResponse struct {
	// in: body
	Body struct {
		// swagger:allOf
		BaseResponse

		Users []User `json:"users"`
	}
}

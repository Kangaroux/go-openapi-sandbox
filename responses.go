package api

// swagger:response baseResponse
type BaseResponse struct {
	// An error message describing what went wrong
	// required: true
	Error string `json:"error"`

	// Whether the request was successful or not
	// required: true
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

func InternalErrorResponse() BaseResponse {
	return ErrorResponse("an internal server error occurred")
}

func ErrorResponse(msg string) BaseResponse {
	return BaseResponse{Error: msg}
}

func OKResponse() BaseResponse {
	return BaseResponse{OK: true}
}

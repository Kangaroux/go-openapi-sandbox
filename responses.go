package api

type BaseResponse struct {
	// An error message describing what went wrong
	// required: true
	Error string `json:"error,omitempty"`

	// Whether the request was successful or not
	// required: true
	OK bool `json:"ok"`
}

type FieldErrorResponse struct {
	// swagger:allOf
	BaseResponse

	// A mapping of field names and errors
	// required: true
	// type: object
	FieldErrors interface{} `json:"field_errors"`
}

type UserResponse struct {
	// swagger:allOf
	BaseResponse

	// The user object
	// required: true
	User User `json:"user"`
}

type ListUserResponse struct {
	// swagger:allOf
	BaseResponse

	// The number of users that were returned
	// required: true
	Count int `json:"count"`

	// A list of user objects
	// required: true
	Users []User `json:"users"`
}

func NewFieldErrorResponse(fieldErrors interface{}, msg string) FieldErrorResponse {
	resp := FieldErrorResponse{}
	resp.Error = msg
	resp.FieldErrors = fieldErrors
	return resp
}

func NewInternalErrorResponse() BaseResponse {
	return NewErrorResponse("an internal server error occurred")
}

func NewErrorResponse(msg string) BaseResponse {
	return BaseResponse{Error: msg}
}

func NewOKResponse() BaseResponse {
	return BaseResponse{OK: true}
}

package api

type BaseResponse struct {
	// An error message describing what went wrong
	// required: true
	Error string `json:"error,omitempty"`

	// Whether the request was successful or not
	// required: true
	OK bool `json:"ok"`
}

// swagger:response baseResponse
type SwaggerBaseResponse struct {
	// in: body
	Body BaseResponse
}

type FieldErrorResponse struct {
	// swagger:allOf
	BaseResponse
	FieldErrors interface{} `json:"field_errors"`
}

// swagger:response fieldErrorResponse
type SwaggerFieldErrorResponse struct {
	// in: body
	Body FieldErrorResponse
}

type UserResponse struct {
	// swagger:allOf
	BaseResponse
	User *User `json:"user,omitempty"`
}

// swagger:response userResponse
type SwaggerUserResponse struct {
	// in: body
	Body UserResponse
}

type ListUserResponse struct {
	// swagger:allOf
	BaseResponse
	Count int     `json:"count"`
	Users []*User `json:"users"`
}

// swagger:response listUserResponse
type SwaggerListUserResponse struct {
	// in: body
	Body ListUserResponse
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

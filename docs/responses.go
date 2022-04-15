package docs

import api "github.com/kangaroux/go-openapi-sandbox"

// swagger:response baseResponse
type SwaggerBaseResponse struct {
	// in: body
	Body api.BaseResponse
}

// swagger:response fieldErrorResponse
type SwaggerFieldErrorResponse struct {
	// in: body
	Body api.FieldErrorResponse
}

// swagger:response userResponse
type SwaggerUserResponse struct {
	// in: body
	Body api.UserResponse
}

// swagger:response listUserResponse
type SwaggerListUserResponse struct {
	// in: body
	Body api.ListUserResponse
}

package docs

import "github.com/pdrum/swagger-automation/api"

// swagger:route POST / 
// Head does some amazing stuff.
// responses:
//   200: headResponse

// This text will appear as description of your response body.
// swagger:response headResponse
type headResponseWrapper struct {
	// in:body
	Body api.FooBarResponse
}

// swagger:parameters idOfHeadEndpoint
type headParamsWrapper struct {
	// This text will appear as description of your request body.
	// in:body
	Body api.HeadRequest
}

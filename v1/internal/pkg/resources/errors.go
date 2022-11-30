package resources

import "errors"

var (
	NoDefaultResponse      = errors.New("no default response set for ResponseWriter")
	NoDefaultErrorResponse = errors.New("no default error response set for ResponseWriter")
)

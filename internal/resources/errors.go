package resources

import "errors"

var (
	NoDefaultResponse         = errors.New("no default response set")
	NoDefaultErrorResponse    = errors.New("no default error response set")
	NoWaggyEntryPointProvided = errors.New("no entrypoint provided to serve")
)

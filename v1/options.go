package v1

import "strings"

// WithDefaultResponse allows you to set a default response to
// be reused by calling waggy.ResponseWriter.WriteDefaultResponse()
func WithDefaultResponse(statusCode int, body []byte) RouteOption {
	r := RouteOption{
		pathParams: nil,
		status:     statusCode,
		body:       body,
	}

	return r
}

// WithDefaultErrorResponse allows you to set a default error response to
// be reused by calling waggy.ResponseWriter.WriteDefaultErrorResponse()
func WithDefaultErrorResponse(statusCode int, err error) RouteOption {
	r := RouteOption{
		pathParams: nil,
		status:     statusCode,
		body:       []byte(err.Error()),
	}

	return r
}

// WithPathParams allows you to use path parameters in the incoming
// request's path. It works by passing in a string representation of
// the request's path with each path parameter surrounded by {}.
// To access each parameter, simply access its values in the map
// stored at waggy.Request.URL.Params. If this option is not provided
// whenever calling waggy.Init(), attempting to retrieve a path parameter
// will return an empty string.
//
// example route: /hello/{name}
//
// To access the value at "name", call r.Url.Params["name"]
func WithPathParams(path string) RouteOption {
	splitPath := strings.Split(path, "/")

	pathParams := make(map[int]string)

	for pathIndex, pathSection := range splitPath {
		pathParams[pathIndex] = pathSection
	}

	opt := RouteOption{
		pathParams: pathParams,
	}

	return opt
}

// RouteOption is used to pass option information to
// the *Request and *ResponseWriter (whichever
// is appropriate) returned from calling waggy.Init()
type RouteOption struct {
	pathParams map[int]string
	status     int
	body       []byte
}

package waggy

import "strings"

func WithDefaultResponse(statusCode int, body []byte) RouteOption {
	r := RouteOption{
		pathParams: nil,
		status:     statusCode,
		body:       body,
	}

	return r
}

func WithDefaultErrorResponse(statusCode int, err error) RouteOption {
	r := RouteOption{
		pathParams: nil,
		status:     statusCode,
		body:       []byte(err.Error()),
	}

	return r
}

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

type RouteOption struct {
	pathParams map[int]string
	status     int
	body       []byte
}

package waggy

import (
	"reflect"
	"runtime"
	"strings"
)

// Init initializes the request and provides a *ResponseWriter
// to use for writing responses, and a *Request to use for
// retrieving info about the incoming HTTP request
func Init(opts ...func() RouteOption) (*ResponseWriter, *Request) {
	pathParamOpt := RouteOption{}

	for _, v := range opts {
		switch getFunctionName(v) {
		case "WithPathParams":
			pathParamOpt = v()
		}
	}

	return Resp(), Req(pathParamOpt)
}

func getFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
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
}

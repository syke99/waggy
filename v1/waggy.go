package v1

import (
	"reflect"
	"runtime"
)

// Init initializes the request and provides a *ResponseWriter
// to use for writing responses, and a *Request to use for
// retrieving info about the incoming HTTP request
func Init(opts ...func() RouteOption) (*ResponseWriter, *Request) {
	reqOpts := make([]RouteOption, 1)
	respOpts := make([]RouteOption, 2)

	for _, v := range opts {
		switch getFunctionName(v) {
		case "WithPathParams":
			reqOpts[0] = v()
		case "WithDefaultResponse":
			respOpts[0] = v()
		case "WithDefaultErrorResponse":
			respOpts[1] = v()
		}
	}

	return Resp(respOpts...), Req(reqOpts...)
}

func getFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

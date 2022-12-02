package v2

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/syke99/waggy/v2/internal/resources"
)

// WaggyHandler is used to handling various http.HandlerFuncs
// mapped by HTTP methods for an individual route
type WaggyHandler struct {
	route          string
	defResp        []byte
	defErrResp     WaggyError
	defErrRespCode int
	handlerMap     map[string]http.HandlerFunc
}

// InitHandler initialized a new WaggyHandler and returns
// a pointer to it
func InitHandler() *WaggyHandler {
	w := WaggyHandler{
		route:          "",
		defResp:        make([]byte, 0),
		defErrResp:     WaggyError{},
		defErrRespCode: 0,
		handlerMap:     make(map[string]http.HandlerFunc),
	}

	return &w
}

// InitHandlerWithRoute initialized a new WaggyHandler with the provided
// route and returns a pointer to it. It is intended to be used whenever
// only compiling an individual *WaggyHandler instead of a full *WaggyRouter
func InitHandlerWithRoute(route string) *WaggyHandler {
	w := WaggyHandler{
		route:          route,
		defResp:        make([]byte, 0),
		defErrResp:     WaggyError{},
		defErrRespCode: 0,
		handlerMap:     make(map[string]http.HandlerFunc),
	}

	return &w
}

// WithDefaultResponse allows you to set a default response for
// individual handlers
func (wh *WaggyHandler) WithDefaultResponse(body []byte) *WaggyHandler {
	wh.defResp = body

	return wh
}

// WithDefaultErrorResponse allows you to set a default error response for
// individual handlers
func (wh *WaggyHandler) WithDefaultErrorResponse(err WaggyError, statusCode int) *WaggyHandler {
	wh.defErrResp = err
	wh.defErrRespCode = statusCode

	return wh
}

// MethodHandler allows you to map a different handler to each HTTP Method
// for a single route.
func (wh *WaggyHandler) MethodHandler(method string, handler http.HandlerFunc) *WaggyHandler {
	wh.handlerMap[method] = handler

	return wh
}

// ServeHTTP serves the route
func (wh *WaggyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	splitRoute := strings.Split(wh.route[1:], "/")

	matchedRoute := os.Getenv(resources.XMatchedRoute.String())[1:]

	splitRequestRoute := strings.Split(matchedRoute, "/")

	vars := make(map[string]string)

	for i, section := range splitRoute {
		beginning := section[:1]
		middle := section[1 : len(section)-1]
		end := section[len(section)-1:]
		if beginning == "{" &&
			end == "}" {
			vars[middle] = splitRequestRoute[i]
		}
	}

	ctx := r.Context()

	if len(wh.defResp) != 0 {
		ctx = context.WithValue(ctx, resources.DefResp, func(w http.ResponseWriter) (int, error) {
			w.Header().Set("Content-Type", http.DetectContentType(wh.defResp))

			return fmt.Fprintln(w, string(wh.defResp))
		})
	}

	if wh.defErrResp.Detail != "" {

		errBytes, _ := json.Marshal(wh.defErrResp)

		w.Header().Set("Content-Type", "application/problem+json")

		ctx = context.WithValue(ctx, resources.DefErr, func(w http.ResponseWriter) {
			fmt.Fprintln(w, string(errBytes))
		})
	}

	if len(vars) != 0 {
		ctx = context.WithValue(ctx, resources.PathParams, vars)
	}

	r = r.Clone(ctx)

	method := os.Getenv(resources.RequestMethod.String())

	wh.handlerMap[method](w, r)
}

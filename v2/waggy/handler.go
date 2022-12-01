package waggy

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/syke99/waggy/v2/waggy/internal/resources"
	"net/http"
	"os"
	"strings"
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
	splitRoute := strings.Split(wh.route, "/")

	splitRequestRoute := strings.Split(os.Getenv(resources.XMatchedRoute.String()), "/")

	vars := make(map[string]string)

	for i, section := range splitRoute {
		if section[:1] == "{" &&
			section[len(section)-1:] == "}" {
			vars[section[1:len(section)-1]] = splitRequestRoute[i]
		}
	}

	ctx := r.Context()

	if len(wh.defResp) != 0 {
		ctx = context.WithValue(r.Context(), resources.DefResp, func(w http.ResponseWriter) (int, error) {
			w.Header().Set("Content-Type", http.DetectContentType(wh.defResp))

			return fmt.Fprintln(w, wh.defResp)
		})
	}

	if wh.defErrResp.Detail != "" {

		errBytes, _ := json.Marshal(wh.defErrResp)

		ctx = context.WithValue(ctx, resources.DefErr, func(w http.ResponseWriter) {
			fmt.Println(errBytes)
		})
	}

	r.WithContext(context.WithValue(ctx, resources.PathParams, vars))

	wh.handlerMap[os.Getenv(resources.RequestMethod.String())](w, r)
}

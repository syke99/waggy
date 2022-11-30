package waggy

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

type WaggyHandler struct {
	route          string
	defResp        []byte
	defErrResp     WaggyError
	defErrRespCode int
	handlerMap     map[string]http.HandlerFunc
}

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

func (wh *WaggyHandler) WithDefaultResponse(body []byte) *WaggyHandler {
	wh.defResp = body

	return wh
}

func (wh *WaggyHandler) WithDefaultErrorResponse(err WaggyError, statusCode int) *WaggyHandler {
	wh.defErrResp = err
	wh.defErrRespCode = statusCode

	return wh
}

func (wh *WaggyHandler) MethodHandler(method string, handler http.HandlerFunc) *WaggyHandler {
	wh.handlerMap[method] = handler

	return wh
}

func (wh *WaggyHandler) ServeHttp(w http.ResponseWriter, r *http.Request) {
	splitRoute := strings.Split(wh.route, "/")

	splitRequestRoute := strings.Split(os.Getenv("X_MATCHED_ROUTE"), "/")

	vars := make(map[string]string)

	for i, section := range splitRoute {
		if section[:1] == "{" &&
			section[len(section)-1:] == "}" {
			vars[section[1:len(section)-1]] = splitRequestRoute[i]
		}
	}

	ctx := r.Context()

	if len(wh.defResp) != 0 {
		ctx = context.WithValue(r.Context(), defResp, func(w http.ResponseWriter) (int, error) {
			w.Header().Set("Content-Type", http.DetectContentType(wh.defResp))

			return fmt.Fprintln(w, wh.defResp)
		})
	}

	if wh.defErrResp.Detail != "" {

		errBytes, _ := json.Marshal(wh.defErrResp)

		ctx = context.WithValue(ctx, defErr, func(w http.ResponseWriter) {
			fmt.Println(errBytes)
		})
	}

	r.WithContext(context.WithValue(ctx, pathParams, vars))

	wh.handlerMap[os.Getenv("REQUEST_METHOD")](w, r)
}

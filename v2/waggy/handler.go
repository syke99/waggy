package waggy

import (
	"context"
	"net/http"
	"os"
	"strings"
)

type WaggyHandler struct {
	route          string
	defResp        []byte
	defErrResp     error
	defErrRespCode int
	handlerMap     map[string]http.HandlerFunc
}

func (wh *WaggyHandler) WithDefaultResponse(body []byte) *WaggyHandler {
	wh.defResp = body

	return wh
}

func (wh *WaggyHandler) WithDefaultErrorResponse(err error, statusCode int) *WaggyHandler {
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

			return w.Write(wh.defResp)
		})
	}

	if wh.defErrResp != nil {
		ctx = context.WithValue(ctx, defErr, func(w http.ResponseWriter) {
			http.Error(w, wh.defErrResp.Error(), wh.defErrRespCode)
		})
	}

	r.WithContext(context.WithValue(ctx, pathParams, vars))

	wh.handlerMap[os.Getenv("REQUEST_METHOD")](w, r)
}

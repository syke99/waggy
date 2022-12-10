package waggy

import (
	"context"
	"fmt"
	"github.com/syke99/waggy/internal/json"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/syke99/waggy/internal/resources"
)

// WaggyHandler is used to handling various http.HandlerFuncs
// mapped by HTTP methods for an individual route
type WaggyHandler struct {
	route                string
	defResp              []byte
	defRespContType      string
	defErrResp           WaggyError
	defErrRespCode       int
	handlerMap           map[string]http.HandlerFunc
	restrictedMethods    map[string]struct{}
	restrictedMethodFunc http.HandlerFunc
	logger               *Logger
	parentLogger         *Logger
	parentLoggerOverride bool
	fullCGI              bool
}

// InitHandler initialized a new WaggyHandler and returns
// a pointer to it
func InitHandler(cgi *FullCGI) *WaggyHandler {
	var o bool
	var err error

	if cgi != nil {
		o, err = strconv.ParseBool(string(*cgi))
		if err != nil {
			o = false
		}
	}

	w := WaggyHandler{
		route:             "",
		defResp:           make([]byte, 0),
		defRespContType:   "",
		defErrResp:        WaggyError{},
		defErrRespCode:    0,
		handlerMap:        make(map[string]http.HandlerFunc),
		restrictedMethods: make(map[string]struct{}),
		logger:            nil,
		parentLogger:      nil,
		fullCGI:           o,
	}

	return &w
}

// InitHandlerWithRoute initialized a new WaggyHandler with the provided
// route and returns a pointer to it. It is intended to be used whenever
// only compiling an individual *WaggyHandler instead of a full *WaggyRouter
func InitHandlerWithRoute(route string, cgi *FullCGI) *WaggyHandler {
	if route[:1] == "/" {
		route = route[1:]
	}
	var o bool
	var err error

	if cgi != nil {
		o, err = strconv.ParseBool(string(*cgi))
		if err != nil {
			o = false
		}
	}

	w := WaggyHandler{
		route:          route,
		defResp:        make([]byte, 0),
		defErrResp:     WaggyError{},
		defErrRespCode: 0,
		handlerMap:     make(map[string]http.HandlerFunc),
		logger:         nil,
		parentLogger:   nil,
		fullCGI:        o,
	}

	return &w
}

// Logger returns the WaggyHandler's Logger. If no parent logger is
// inherited from a WaggyRouter, or you provided a OverrideParentLogger
// whenever adding a Logger to the WaggyHandler, then the WaggyHandler's
// Logger will be returned. If no logger has been set, then this method
// will return nil
func (wh *WaggyHandler) Logger() *Logger {
	if wh.parentLogger == nil ||
		(wh.logger != nil && wh.parentLoggerOverride) {
		return wh.logger
	}
	return wh.parentLogger
}

// WithLogger allows you to set a logger for wh
func (wh *WaggyHandler) WithLogger(logger *Logger, parentOverride ParentLoggerOverrider) *WaggyHandler {
	wh.logger = logger
	if parentOverride == nil {
		wh.parentLoggerOverride = false
		return wh
	}

	wh.parentLoggerOverride = parentOverride()

	return wh
}

// WithDefaultLogger sets wh's logger to the default Logger
func (wh *WaggyHandler) WithDefaultLogger() *WaggyHandler {
	l := Logger{
		logLevel: Info.level(),
		key:      "",
		message:  "",
		err:      "",
		vals:     make(map[string]interface{}),
		log:      os.Stderr,
	}

	wh.logger = &l

	return wh
}

func (wh *WaggyHandler) inheritLogger(lp *Logger) {
	wh.parentLogger = lp
}

func (wh *WaggyHandler) inheritFullCgiFlag(cgi bool) {
	wh.fullCGI = cgi
}

// WithDefaultResponse allows you to set a default response for
// individual handlers
func (wh *WaggyHandler) WithDefaultResponse(contentType string, body []byte) *WaggyHandler {
	wh.defResp = body
	wh.defRespContType = contentType

	return wh
}

// WithDefaultErrorResponse allows you to set a default error response for
// individual handlers
func (wh *WaggyHandler) WithDefaultErrorResponse(err WaggyError, statusCode int) *WaggyHandler {
	wh.defErrResp = err
	wh.defErrRespCode = statusCode

	return wh
}

// AllHTTPMethods allows you to easily set a
// handler all HTTP Methods
var AllHTTPMethods = func() string {
	return "ALL"
}

// MethodHandler allows you to map a different handler to each HTTP Method
// for a single route.
func (wh *WaggyHandler) MethodHandler(method string, handler http.HandlerFunc) *WaggyHandler {
	if _, ok := resources.AllHTTPMethods()[method]; !ok {
		return wh
	}

	if method == "ALL" {
		for k, _ := range resources.AllHTTPMethods() {
			wh.handlerMap[k] = handler
		}
	} else {
		wh.handlerMap[method] = handler
	}

	return wh
}

// RestrictMethods is a variadic function for restricting a handler from being able
// to be executed on the given methods
func (wh *WaggyHandler) RestrictMethods(methods ...string) *WaggyHandler {
	for _, method := range methods {
		if _, ok := resources.AllHTTPMethods()[method]; !ok {
			continue
		}
		wh.restrictedMethods[method] = struct{}{}
	}

	return wh
}

// WithRestrictedMethodHandler allows you to set an http.HandlerFunc to be used
// whenever a request with a restricted HTTP Method is hit. Whenever ServeHTTP is
// called, if this method has not been called and a restricted method has been set
// and is hit by the incoming request, it will return a generic 405 error, instead
func (wh *WaggyHandler) WithRestrictedMethodHandler(fn http.HandlerFunc) *WaggyHandler {
	wh.restrictedMethodFunc = fn

	return wh
}

// ServeHTTP serves the route
func (wh *WaggyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if _, ok := wh.restrictedMethods[r.Method]; ok {
		if wh.restrictedMethodFunc != nil {
			wh.restrictedMethodFunc(w, r)
			return
		}

		r.URL.Opaque = ""
		rRoute := r.URL.Path

		methodNotAllowed := WaggyError{
			Title:    "Method Not Allowed",
			Detail:   "method not allowed",
			Status:   405,
			Instance: rRoute,
		}

		wh.defErrResp = methodNotAllowed

		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Header().Set("Content-Type", "application/problem+json")
		fmt.Fprintln(w, json.BuildJSONStringFromWaggyError(wh.defErrResp.Type, wh.defErrResp.Title, wh.defErrResp.Detail, wh.defErrResp.Status, wh.defErrResp.Instance, wh.defErrResp.Field))
		return
	}

	if len(wh.route) != 1 && wh.route[:1] == "/" {
		wh.route = wh.route[1:]
	}

	if rr := r.Context().Value(resources.RootRoute); rr != nil {
		ctx := context.WithValue(r.Context(), resources.MatchedRoute, "/")

		r = r.Clone(ctx)
	}

	splitRoute := strings.Split(wh.route, "/")

	vars := make(map[string]string)

	route := ""

	if mr := r.Context().Value(resources.MatchedRoute); mr != nil {
		route = r.Context().Value(resources.MatchedRoute).(string)
	} else {
		r.URL.Opaque = ""

		route = r.URL.Path
	}

	splitRequestRoute := []string{"/"}

	if route != "/" {
		splitRequestRoute = strings.Split(route, "/")

		if route[:1] == "/" {
			splitRequestRoute = strings.Split(route[1:], "/")
		}
	}

	for i, section := range splitRoute {
		if section == "" || section == "/" {
			continue
		}

		beginning := section[:1]
		middle := section[1 : len(section)-1]
		end := section[len(section)-1:]
		if beginning == "{" &&
			end == "}" {
			vars[middle] = splitRequestRoute[i]
		}
	}

	ctx := r.Context()

	queryParams := make(map[string][]string)

	q, _ := url.QueryUnescape(r.URL.RawQuery)

	qp := strings.Split(q, "&")

	if !wh.fullCGI {
		qp = os.Args[1:]
	}

	if len(qp) != 0 && qp[0] != "" {
		for _, _qp := range qp {
			sqp := strings.Split(_qp, "=")

			key := ""
			value := ""

			if len(sqp) == 2 {
				key = sqp[0]
				value = sqp[1]

				queryParams[key] = append(queryParams[key], value)
			}
		}
	}

	if len(wh.defResp) != 0 {
		ctx = context.WithValue(ctx, resources.DefResp, func(w http.ResponseWriter) {
			w.Header().Set("Content-Type", wh.defRespContType)

			fmt.Fprintln(w, string(wh.defResp))
		})
	}

	if wh.defErrResp.Detail != "" {
		w.Header().Set("Content-Type", "application/problem+json")

		ctx = context.WithValue(ctx, resources.DefErr, func(w http.ResponseWriter) {
			fmt.Fprintln(w, json.BuildJSONStringFromWaggyError(wh.defErrResp.Type, wh.defErrResp.Title, wh.defErrResp.Detail, wh.defErrResp.Status, wh.defErrResp.Instance, wh.defErrResp.Field))
		})
	}

	if len(vars) != 0 {
		ctx = context.WithValue(ctx, resources.PathParams, vars)
	}

	if len(queryParams) != 0 {
		ctx = context.WithValue(ctx, resources.QueryParams, queryParams)
	}

	if wh.logger != nil {
		ctx = context.WithValue(ctx, resources.Logger, wh.logger)
	}

	r = r.Clone(ctx)

	wh.handlerMap[r.Method](w, r)
}

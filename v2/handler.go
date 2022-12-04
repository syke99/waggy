package v2

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/syke99/waggy/v2/internal/resources"
)

// WaggyHandler is used to handling various http.HandlerFuncs
// mapped by HTTP methods for an individual route
type WaggyHandler struct {
	route                string
	defResp              []byte
	defErrResp           WaggyError
	defErrRespCode       int
	handlerMap           map[string]http.HandlerFunc
	logger               *Logger
	parentLogger         *Logger
	parentLoggerOverride bool
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
		logger:         nil,
		parentLogger:   nil,
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

type FileServer func(w http.ResponseWriter, r *http.Request)

func (wh *WaggyHandler) FileServer(filePath string) (FileServer, error) {
	if filePath == "" {
		return nil, errors.New("no path to file provided")
	}

	return func(w http.ResponseWriter, r *http.Request) {
		file, err := os.Open(filePath)
		if err != nil {
			wh.Logger().
				Err(err).
				Msg("Method:", "ServeFile").
				Log()

			w.WriteHeader(http.StatusInternalServerError)
			w.Header().Set("content-type", "text/plain")
			w.Write([]byte("Error opening file"))
		}

		cTBuf := make([]byte, 0, 512)

		_, err = file.Read(cTBuf)
		if err != nil {
			wh.Logger().
				Err(err).
				Msg("Method:", "ServeFile").
				Log()

			w.WriteHeader(http.StatusInternalServerError)
			w.Header().Set("content-type", "text/plain")
			w.Write([]byte("Error reading file"))
		}

		cTType := http.DetectContentType(cTBuf)

		w.Header().Set("content-type", cTType)

		if _, err = io.Copy(w, file); err != nil {
			wh.Logger().
				Err(err).
				Msg("Method:", "ServeFile").
				Log()

			w.WriteHeader(http.StatusInternalServerError)
			w.Header().Set("content-type", "text/plain")
			w.Write([]byte("Error returning file"))
		}
	}, nil
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

		ctx = context.WithValue(ctx, resources.DefErr, func(w http.ResponseWriter) (int, error) {
			return fmt.Fprintln(w, string(errBytes))
		})
	}

	if len(vars) != 0 {
		ctx = context.WithValue(ctx, resources.PathParams, vars)
	}

	r = r.Clone(ctx)

	method := os.Getenv(resources.RequestMethod.String())

	wh.handlerMap[method](w, r)
}

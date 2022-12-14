package waggy

import (
	"context"
	"fmt"
	"github.com/syke99/waggy/internal/json"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/syke99/waggy/internal/resources"
)

// Router is used for routing incoming HTTP requests to
// specific *Handlers by the route provided whenever you call
// Handle on the return router and provide a route for the *Handler
// you provide
type Router struct {
	logger      *Logger
	router      map[string]*Handler
	noRoute     WaggyError
	noRouteFunc http.HandlerFunc
	FullServer  bool
}

// InitRouter initializes a new Router and returns a pointer
// to it
func InitRouter(cgi *FullServer) *Router {
	var o bool
	var err error

	if cgi != nil {
		o, err = strconv.ParseBool(string(*cgi))
		if err != nil {
			o = false
		}
	}

	r := Router{
		logger: nil,
		router: make(map[string]*Handler),
		noRoute: WaggyError{
			Title:    "Resource not found",
			Detail:   "route not found",
			Status:   404,
			Instance: "/",
		},
		FullServer: o,
	}

	return &r
}

// Handle allows you to map a *Handler for a specific route. Just
// in the popular gorilla/mux router, you can specify path parameters
// by wrapping them with {} and they can later be accessed by calling
// Vars(r)
func (wr *Router) Handle(route string, handler *Handler) *Router {
	handler.route = route
	handler.inheritLogger(wr.logger)
	handler.inheritFullServerFlag(wr.FullServer)
	wr.router[route] = handler

	return wr
}

func (wr *Router) Routes() []string {
	r := make([]string, 0)

	for k, _ := range wr.router {
		r = append(r, k)
	}

	return r
}

// WithLogger allows you to set a Logger for the entire router. Whenever
// Handle is called, this logger will be passed to the *Handler
// being handled for the given route.
func (wr *Router) WithLogger(logger *Logger) *Router {
	wr.logger = logger

	return wr
}

// WithDefaultLogger sets wr's logger to the default Logger
func (wr *Router) WithDefaultLogger() *Router {
	l := Logger{
		logLevel: Info.level(),
		key:      "",
		message:  "",
		err:      "",
		vals:     make(map[string]interface{}),
		log:      os.Stderr,
	}

	wr.logger = &l

	return wr
}

// WithNoRouteHandler allows you to set an http.HandlerFunc to be used whenever
// no route is found. If this method is not called and the ServeHTTP method
// has been called, then it will return a generic 404 response, instead
func (wr *Router) WithNoRouteHandler(fn http.HandlerFunc) *Router {
	wr.noRouteFunc = fn

	return wr
}

// Logger returns the Router's logger
func (wr *Router) Logger() *Logger {
	return wr.logger
}

// ServeHTTP satisfies the http.Handler interface and calls the stored
// handler at the route of the incoming HTTP request
func (wr *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rt := ""

	r.URL.Opaque = ""
	rRoute := r.URL.Path

	if rRoute == "" || rRoute == "/" {
		if handler, ok := wr.router["/"]; !ok {
			w.WriteHeader(http.StatusMethodNotAllowed)
			wr.noRouteResponse(w, r)
			return
		} else {
			ctx := context.WithValue(r.Context(), resources.RootRoute, true)

			r = r.Clone(ctx)

			handler.ServeHTTP(w, r)
			return
		}
	}

	for key, handler := range wr.router {
		if key == "/" {
			continue
		}

		if rRoute[:1] == "/" {
			rRoute = rRoute[1:]
		}

		splitRoute := strings.Split(rRoute, "/")

		if key[:1] == "/" {
			key = key[1:]
		}

		splitKey := strings.Split(key, "/")

		for i, section := range splitKey {
			if len(section) == 0 {
				continue
			}

			beginning := section[:1]
			end := section[len(section)-1:]

			// check if this section is a query param
			if (beginning == "{" &&
				end == "}") && (len(splitRoute) != len(splitKey)) {
				continue
			}

			if (beginning == "{" &&
				end == "}") && (len(splitRoute) == len(splitKey)) {
				rt = key
			}

			// if the route sections don't match and aren't query
			// params, break out as these are not the correctly matched
			// routes
			if i > len(splitRoute) || splitRoute[i] != section && rt == "" {
				break
			}

			if len(splitKey) > len(splitRoute) &&
				i == len(splitRoute) &&
				rt == "" {
				rt = key
			}

			// if the end of splitRoute is reached, and we haven't
			// broken out of the loop to move on to the next route,
			// then the routes match
			if (i == len(splitKey)-1 || i == len(splitRoute)) &&
				rt == "" {
				rt = key
			}
		}

		if rt != "" {
			ctx := context.WithValue(r.Context(), resources.MatchedRoute, rRoute)

			r = r.Clone(ctx)

			handler.ServeHTTP(w, r)
			return
		}
	}

	wr.noRouteResponse(w, r)
	return
}

func (wr *Router) noRouteResponse(w http.ResponseWriter, r *http.Request) {
	if wr.noRouteFunc != nil {
		wr.noRouteFunc(w, r)
		return
	}

	w.WriteHeader(http.StatusNotFound)
	w.Header().Set("Content-Type", "application/problem+json")
	fmt.Fprintln(w, json.BuildJSONStringFromWaggyError(wr.noRoute.Type, wr.noRoute.Title, wr.noRoute.Detail, wr.noRoute.Status, wr.noRoute.Instance, wr.noRoute.Field))
}

package waggy

import (
	"net/http"
	"os"
	"strconv"

	"github.com/syke99/waggy/internal/resources"
)

// WaggyRouter is used for routing incoming HTTP requests to
// specific *WaggyHandlers by the route provided whenever you call
// Handle on the return router and provide a route for the *WaggyHandler
// you provide
type WaggyRouter struct {
	logger  *Logger
	router  map[string]*WaggyHandler
	fullCGI bool
}

// InitRouter initializes a new WaggyRouter and returns a pointer
// to it
func InitRouter(cgi *FullCGI) *WaggyRouter {
	var o bool
	var err error

	if cgi != nil {
		o, err = strconv.ParseBool(string(*cgi))
		if err != nil {
			o = false
		}
	}

	r := WaggyRouter{
		router:  make(map[string]*WaggyHandler),
		fullCGI: o,
	}

	return &r
}

// Handle allows you to map a *WaggyHandler for a specific route. Just
// in the popular gorilla/mux router, you can specify path parameters
// by wrapping them with {} and they can later be accessed by calling
// Vars(r)
func (wr *WaggyRouter) Handle(route string, handler *WaggyHandler) *WaggyRouter {
	handler.route = route
	handler.inheritLogger(wr.logger)
	wr.router[route] = handler

	return wr
}

// WithLogger allows you to set a Logger for the entire router. Whenever
// Handle is called, this logger will be passed to the *WaggyHandler
// being handled for the given route.
func (wr *WaggyRouter) WithLogger(logger *Logger) *WaggyRouter {
	wr.logger = logger

	return wr
}

// WithDefaultLogger sets wr's logger to the default Logger
func (wr *WaggyRouter) WithDefaultLogger() *WaggyRouter {
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

// Logger returns the WaggyRouter's logger
func (wr *WaggyRouter) Logger() *Logger {
	return wr.logger
}

// ServeHTTP satisfies the http.Handler interface and calls the stored
// handler at the route of the incoming HTTP request
func (wr *WaggyRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	wr.router[os.Getenv(resources.XMatchedRoute.String())].ServeHTTP(w, r)
}

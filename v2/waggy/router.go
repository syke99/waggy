package waggy

import (
	"net/http"
	"os"
)

// WaggyRouter is used for routing incoming HTTP requests to
// specific *WaggyHandlers by the route provided whenever you call
// Handle on the return router and provide a route for the *WaggyHandler
// you provide
type WaggyRouter struct {
	router map[string]*WaggyHandler
}

// InitRouter initializes a new WaggyRouter and returns a pointer
// to it
func InitRouter() *WaggyRouter {
	r := WaggyRouter{
		router: make(map[string]*WaggyHandler),
	}

	return &r
}

// Handle allows you to map a *WaggyHandler for a specific route. Just
// in the popular gorilla/mux router, you can specify path parameters
// by wrapping them with {} and they can later be accessed by calling
// waggy.Vars(r)
func (wr *WaggyRouter) Handle(route string, handler *WaggyHandler) {
	handler.route = route
	wr.router[route] = handler
}

// ServeHTTP satisfies the http.Handler interface and calls the stored
// handler at the route of the incoming HTTP request
func (wr *WaggyRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	wr.router[os.Getenv("X_MATCHED_ROUTE")].ServeHttp(w, r)
}

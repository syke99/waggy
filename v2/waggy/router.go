package waggy

import (
	"net/http"
	"os"
)

type WaggyRouter struct {
	router map[string]*WaggyHandler
}

func InitRouter() *WaggyRouter {
	r := WaggyRouter{
		router: make(map[string]*WaggyHandler),
	}

	return &r
}

func (wr *WaggyRouter) Handle(route string, handler *WaggyHandler) {
	handler.route = route
	wr.router[route] = handler
}

func (wr *WaggyRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	wr.router[os.Getenv("X_MATCHED_ROUTE")].ServeHttp(w, r)
}

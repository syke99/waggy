package resources

import "net/http"

type ContextKey int

const (
	DefResp ContextKey = iota
	DefErr
	MatchedRoute
	RootRoute
	PathParams
	QueryParams
	Logger
)

var AllHTTPMethods = func() []string {
	return []string{
		http.MethodGet,
		http.MethodDelete,
		http.MethodPut,
		http.MethodPost,
		http.MethodPatch,
		http.MethodConnect,
		http.MethodTrace,
		http.MethodHead,
	}
}

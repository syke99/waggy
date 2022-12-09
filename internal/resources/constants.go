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

var AllHTTPMethods = func() map[string]struct{} {
	m := make(map[string]struct{})
	m[http.MethodGet] = struct{}{}
	m[http.MethodPut] = struct{}{}
	m[http.MethodPost] = struct{}{}
	m[http.MethodPatch] = struct{}{}
	m[http.MethodOptions] = struct{}{}
	m[http.MethodConnect] = struct{}{}
	m[http.MethodDelete] = struct{}{}
	m[http.MethodTrace] = struct{}{}
	m[http.MethodHead] = struct{}{}

	return m
}

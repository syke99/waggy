package middleware

import "net/http"

type MiddleWare func(handler http.Handler) http.Handler

func PassThroughMiddleWare(middle []MiddleWare, handler http.HandlerFunc) http.HandlerFunc {
	for _, mw := range middle {
		handler = mw(handler).ServeHTTP
	}

	return handler
}

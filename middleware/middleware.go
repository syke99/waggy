package middleware

import "net/http"

type MiddleWare func(handler http.Handler)

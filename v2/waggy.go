package v2

import (
	"fmt"
	"net/http"
	"net/http/cgi"
	"reflect"

	"github.com/syke99/waggy/v2/internal/resources"
)

// WaggyEntryPoint is used as a type constraint whenever calling
// Serve so that only a *WaggyRouter or *WaggyHandler can
// be used and not a bare http.Handler
type WaggyEntryPoint interface {
	*WaggyRouter | *WaggyHandler
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

// WriteDefaultResponse returns the result (number of bytes written
// and a nil value, or the error of that write) of writing the set
// default response inside the handler it is being used inside of.
// If no default response has been set, this function will return
// an error.
func WriteDefaultResponse(w http.ResponseWriter, r *http.Request) {
	rv := r.Context().Value(resources.DefResp)
	if rv == nil {
		fmt.Fprintln(w, resources.NoDefaultResponse.Error())
	}

	reflect.ValueOf(rv).Call([]reflect.Value{reflect.ValueOf(w)})
}

// WriteDefaultErrorResponse returns the result of writing the set
// default error response inside the handler it is being used inside of.
// If no default error response has been set, this function will return
// an error.
func WriteDefaultErrorResponse(w http.ResponseWriter, r *http.Request) {
	rv := r.Context().Value(resources.DefErr)
	if rv == nil {
		fmt.Fprintln(w, resources.NoDefaultErrorResponse.Error())
	}

	results := reflect.ValueOf(rv).Call([]reflect.Value{reflect.ValueOf(w)})

	if len(results) != 0 {
		fmt.Fprintln(w, results[0].Interface().(error).Error())
	}
}

// Vars returns the route variables for the current request, if any.
func Vars(r *http.Request) map[string]string {
	if rv := r.Context().Value(resources.PathParams); rv != nil {
		return rv.(map[string]string)
	}
	return nil
}

// Serve wraps a call to cgi.serve and also uses a type constraint of
// WaggyEntryPoint so that only a *WaggyRouter or *WaggyHandler can be
// used in the call to Serve and not accidentally allow calling
// a bare http.Handler
func Serve[W WaggyEntryPoint](entryPoint W) error {
	return cgi.Serve(entryPoint)
}

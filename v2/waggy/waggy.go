package waggy

import (
	"errors"
	"net/http"
	"reflect"

	"github.com/syke99/waggy/v2/waggy/internal/resources"
)

// WriteDefaultResponse returns the result (number of bytes written
// and a nil value, or the error of that write) of writing the set
// default response inside the handler it is being used inside of.
// If no default response has been set, this function will return
// an error.
func WriteDefaultResponse(w http.ResponseWriter, r *http.Request) (int, error) {
	rv := r.Context().Value(resources.DefResp)
	if rv == nil {
		return 0, resources.NoDefaultResponse
	}

	results := reflect.ValueOf(rv).Call([]reflect.Value{reflect.ValueOf(w)})

	return int(results[0].Int()), errors.New(results[1].String())
}

// WriteDefaultErrorResponse returns the result of writing the set
// default error response inside the handler it is being used inside of.
// If no default error response has been set, this function will return
// an error.
func WriteDefaultErrorResponse(w http.ResponseWriter, r *http.Request) error {
	rv := r.Context().Value(resources.DefErr)
	if rv == nil {
		return resources.NoDefaultErrorResponse
	}

	results := reflect.ValueOf(rv).Call([]reflect.Value{reflect.ValueOf(w)})

	return results[0].Interface().(error)
}

// Vars returns the route variables for the current request, if any.
func Vars(r *http.Request) map[string]string {
	if rv := r.Context().Value(resources.PathParams); rv != nil {
		return rv.(map[string]string)
	}
	return nil
}

package waggy

import (
	"errors"
	"net/http"
	"reflect"
)

type contextKey int

const (
	defResp contextKey = iota
	defErr
	pathParams
)

// WriteDefaultResponse returns the result (number of bytes written
// and a nil value, or the error of that write) of writing the set
// default response inside the handler it is being used inside of.
// If no default response has been set, this function will return
// an error.
func WriteDefaultResponse(w http.ResponseWriter, r *http.Request) (int, error) {
	rv := r.Context().Value(defResp)
	if rv == nil {
		return 0, errors.New("no default response set")
	}

	results := reflect.ValueOf(rv).Call([]reflect.Value{reflect.ValueOf(w)})

	return int(results[0].Int()), errors.New(results[1].String())
}

// WriteDefaultErrorResponse returns the result of writing the set
// default error response inside the handler it is being used inside of.
// If no default error response has been set, this function will return
// an error.
func WriteDefaultErrorResponse(w http.ResponseWriter, r *http.Request) error {
	rv := r.Context().Value(defErr)
	if rv == nil {
		return errors.New("no default error response set")
	}

	results := reflect.ValueOf(rv).Call([]reflect.Value{reflect.ValueOf(w)})

	return results[0].Interface().(error)
}

// Vars returns the route variables for the current request, if any.
func Vars(r *http.Request) map[string]string {
	if rv := r.Context().Value(pathParams); rv != nil {
		return rv.(map[string]string)
	}
	return nil
}

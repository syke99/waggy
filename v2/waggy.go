package v2

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/syke99/waggy/v2/internal/resources"
)

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

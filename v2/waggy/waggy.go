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

func WriteDefaultResponse(w http.ResponseWriter, r *http.Request) (int, error) {
	rv := r.Context().Value(defResp)
	if rv == nil {
		return 0, nil
	}

	results := reflect.ValueOf(rv).Call([]reflect.Value{reflect.ValueOf(w)})

	return int(results[0].Int()), errors.New(results[1].String())
}

func WriteDefaultErrorResponse(w http.ResponseWriter, r *http.Request) {
	rv := r.Context().Value(defErr)
	if rv == nil {
		return
	}

	reflect.ValueOf(rv).Call([]reflect.Value{reflect.ValueOf(w)})
}

func Vars(r *http.Request) map[string]string {
	if rv := r.Context().Value(pathParams); rv != nil {
		return rv.(map[string]string)
	}
	return nil
}

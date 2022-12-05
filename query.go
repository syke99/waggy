package waggy

import (
	"github.com/syke99/waggy/internal/resources"
	"net/http"
)

// QueryParams is a map containing all the query params from the provided *http.Request
type QueryParams = map[string][]string

// Query returns QueryParams from the provided *http.Request
func Query(r *http.Request) QueryParams {
	if rv := r.Context().Value(resources.QueryParams); rv != nil {
		return rv.(QueryParams)
	}
	return nil
}

// Get returns the first value stored in q with the provided key
func (q QueryParams) Get(key string) string {
	v, ok := q[key]
	if !ok {
		return ""
	}

	return v[0]
}

// Set sets the provided val in q with the provided key. If an existing
// value/set of values is already stored at the provided key, then Set
// will override that value
func (q QueryParams) Set(key string, val string) {
	if _, ok := q[key]; ok {
		delete(q, key)
	}

	newSlice := append(make([]string, 0), val)

	q[key] = newSlice
}

// Add either appends the provided val at the provided key stored in q, or,
// if no value is currently stored at the provided key, then a new slice will
// be stored at the provided key and the provided val appended to it
func (q QueryParams) Add(key string, val string) {
	if _, ok := q[key]; !ok {
		q[key] = make([]string, 0)
	}

	q[key] = append(q[key], val)
}

// Del deletes the value(s) stored in q at the provided key
func (q QueryParams) Del(key string) {
	delete(q, key)
}

// Values returns a slice of all values stored in q with the provided key
func (q QueryParams) Values(key string) []string {
	if _, ok := q[key]; !ok {
		return make([]string, 0)
	}

	return q[key]
}

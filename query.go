package waggy

import (
	"github.com/syke99/waggy/internal/resources"
	"net/http"
)

// QueryParams contains a map containing all the query params from the provided *http.Request
type QueryParams struct {
	qp map[string][]string
}

// Query returns QueryParams from the provided *http.Request
func Query(r *http.Request) *QueryParams {
	if rv := r.Context().Value(resources.QueryParams); rv != nil {
		qp := rv.(map[string][]string)

		q := QueryParams{
			qp: qp,
		}

		return &q
	}
	return nil
}

// Get returns the first value stored in q with the provided key
func (q *QueryParams) Get(key string) string {
	v, ok := q.qp[key]
	if !ok {
		return ""
	}

	return v[0]
}

// Set sets the provided val in q with the provided key. If an existing
// value/set of values is already stored at the provided key, then Set
// will override that value
func (q *QueryParams) Set(key string, val string) {
	if _, ok := q.qp[key]; ok {
		delete(q.qp, key)
	}

	newSlice := append(make([]string, 0), val)

	q.qp[key] = newSlice
}

// Add either appends the provided val at the provided key stored in q, or,
// if no value is currently stored at the provided key, then a new slice will
// be stored at the provided key and the provided val appended to it
func (q *QueryParams) Add(key string, val string) {
	if _, ok := q.qp[key]; !ok {
		q.qp[key] = make([]string, 0)
	}

	q.qp[key] = append(q.qp[key], val)
}

// Del deletes the value(s) stored in q at the provided key
func (q *QueryParams) Del(key string) {
	delete(q.qp, key)
}

// Values returns a slice of all values stored in q with the provided key
func (q *QueryParams) Values(key string) []string {
	if _, ok := q.qp[key]; !ok {
		return make([]string, 0)
	}

	return q.qp[key]
}

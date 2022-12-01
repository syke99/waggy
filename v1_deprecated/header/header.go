package header

import (
	"github.com/syke99/waggy/v1/internal/pkg/resources"
	"os"
	"strings"
)

type Header struct {
	headers map[string][]string
}

func GetHeaders() *Header {
	var i resources.Env

	headers := make(map[string][]string)

	for i < 22 {
		if i.String() == resources.ContentType.String() {
			headers[i.String()] = strings.Split(os.Getenv(i.String()), "; ")
			i++
			continue
		}
		headers[i.String()] = strings.Split(os.Getenv(i.String()), ", ")
		i++
	}

	h := Header{headers: headers}

	return &h
}

// CreateHeader is a convenience function for creating a *Header with
// the given key and sets the given values to that key
func CreateHeader(key string, values ...string) *Header {
	h := Header{
		headers: make(map[string][]string),
	}

	for _, value := range values {
		h.Add(key, value)
	}

	return &h
}

// Add appends the value to the slice of strings for the given key
func (h *Header) Add(key string, value string) {
	if h.headers == nil {
		h.headers = make(map[string][]string)
	}

	if _, ok := h.headers[key]; !ok {
		h.headers[key] = make([]string, 0)

		h.headers[key] = append(h.headers[key], value)
		return
	}

	h.headers[key] = append(h.headers[key], value)
}

// Del deletes the value stored in the Header with the given key
func (h *Header) Del(key string) {
	delete(h.headers, key)
}

// Get returns the first value stored in the Header with the given key
func (h *Header) Get(key string) string {
	if len(h.headers) == 0 {
		return ""
	}
	return h.headers[key][0]
}

// Has returns whether a Header value has been set with the given key
func (h *Header) Has(key string) bool {
	_, ok := h.headers[key]

	return ok
}

// Set overrides the values stored at the given key with the given value
func (h *Header) Set(key string, value string) {
	if h.headers == nil {
		h.headers = make(map[string][]string)
	}

	h.headers[key] = make([]string, 0)

	h.headers[key] = append(h.headers[key], value)
}

// Values returns all the values stored with the given key
func (h *Header) Values(key string) []string {
	return h.headers[key]
}

// Loop returns a map[string][]string of Header values to loop over
func (h *Header) Loop() map[string][]string {
	return h.headers
}

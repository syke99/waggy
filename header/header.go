package header

import (
	"github.com/syke99/waggy/internal/pkg/resources"
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
		headers[i.String()] = strings.Split(os.Getenv(i.String()), "; ")
		i++
	}

	h := Header{headers: headers}

	return &h
}

func (h *Header) Add(key string, value string) {
	if _, ok := h.headers[key]; !ok {
		h.headers[key] = make([]string, 0)

		h.headers[key] = append(h.headers[key], value)
		return
	}

	h.headers[key] = append(h.headers[key], value)
}

func (h *Header) Del(key string) {
	delete(h.headers, key)
}

func (h *Header) Get(key string) string {
	if len(h.headers) == 0 {
		return ""
	}
	return h.headers[key][0]
}

func (h *Header) Set(key string, value string) {
	h.headers[key] = make([]string, 0)

	h.headers[key] = append(h.headers[key], value)
}

func (h *Header) Values(key string) []string {
	return h.headers[key]
}

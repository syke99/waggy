package main

import (
	"fmt"
	"github.com/syke99/waggy/header"
	"github.com/syke99/waggy/internal/pkg/resources"
	"net/http"
	"os"
	"strings"
)

type WaggyResponseWriter struct {
	status int
	Header *header.Header
}

func Response() *WaggyResponseWriter {
	h := header.Header{}

	rw := WaggyResponseWriter{Header: &h}

	return &rw
}

func (w *WaggyResponseWriter) Write(body []byte) (int, error) {
	if !w.Header.Has("Content-Type") {
		w.Header.Set("Content-Type", http.DetectContentType(body))
	}

	payload := make([]byte, 0)

	payload = append(payload, []byte(fmt.Sprintf("%s %d %s\n", os.Getenv(resources.Scheme.String()), w.status, resources.GetStatusName(w.status)))...)

	headerLines := make([][]byte, 0)

	for k, v := range w.Header.Loop() {
		if k == resources.ContentType.String() {
			headerLines = append(headerLines, []byte(fmt.Sprintf("%s: %s\n", k, strings.Join(v, ", "))))
		}
	}

	for _, headerLine := range headerLines {
		payload = append(payload, headerLine...)
	}

	payload = append(payload, []byte("\n")...)

	return os.Stdout.Write(payload)
}

func (w *WaggyResponseWriter) WriteHeader(statusCode int) {
	w.status = statusCode
}

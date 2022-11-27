package waggy

import (
	"encoding/json"
	"fmt"
	"github.com/syke99/waggy/header"
	"github.com/syke99/waggy/internal/pkg/models"
	"github.com/syke99/waggy/internal/pkg/resources"
	"net/http"
	"os"
	"strings"
)

type WaggyResponseWriter struct {
	status resources.StatusCode
	Header *header.Header
}

func Response() *WaggyResponseWriter {
	h := header.Header{}

	rw := WaggyResponseWriter{Header: &h}

	return &rw
}

func (w *WaggyResponseWriter) WriteHeader(statusCode int) {
	w.status = resources.StatusCode(statusCode)
}

func (w *WaggyResponseWriter) Write(body []byte) (int, error) {
	if !w.Header.Has("Content-Type") {
		w.Header.Set("Content-Type", http.DetectContentType(body))
	}

	payload := w.buildResponse(body)

	return os.Stdout.Write(payload)
}

func (w *WaggyResponseWriter) Error(statusCode int, error string) (int, error) {

	w.status = resources.StatusCode(statusCode)

	w.Header.Set("Content-Type", "application/problem+json")

	err := models.ErrReponse{
		Type:   os.Getenv(resources.FullURL.String()),
		Detail: error,
		Status: statusCode,
	}

	errBytes, _ := json.Marshal(err)

	payload := w.buildResponse(errBytes)

	return os.Stdout.Write(payload)
}

func (w *WaggyResponseWriter) buildResponse(payload []byte) []byte {

	response := make([]byte, 0)

	response = append(response, []byte(fmt.Sprintf("%s %d %s\n", os.Getenv(resources.Scheme.String()), w.status, w.status.GetStatusName()))...)

	headerLines := make([][]byte, 0)

	for k, v := range w.Header.Loop() {
		if k == "" {
			continue
		}

		if k == resources.ContentType.String() {
			headerLines = append(headerLines, []byte(fmt.Sprintf("%s: %s\n", k, strings.Join(v, "; "))))
		}

		headerLines = append(headerLines, []byte(fmt.Sprintf("%s: %s\n", k, strings.Join(v, ", "))))
	}

	for _, headerLine := range headerLines {
		response = append(response, headerLine...)
	}

	response = append(response, []byte("\n")...)

	response = append(response, payload...)

	return response
}

package waggy

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/syke99/waggy/header"
	"github.com/syke99/waggy/internal/pkg/models"
	"github.com/syke99/waggy/internal/pkg/resources"
	"io"
	"net/http"
	"os"
	"strings"
)

// WaggyResponseWriter used for writing an HTTP Response
type WaggyResponseWriter struct {
	status resources.StatusCode
	Header *header.Header
	writer io.Writer
	buffer *bytes.Buffer
}

// Response initializes a new WaggyResponseWriter to be used to write HTTP Responses
func Response() *WaggyResponseWriter {
	h := header.Header{}

	rw := WaggyResponseWriter{
		Header: &h,
		writer: os.Stdout,
		buffer: bytes.NewBuffer(make([]byte, 0)),
	}

	return &rw
}

// WriteHeader writes the provided statusCode Header
func (w *WaggyResponseWriter) WriteHeader(statusCode int) {
	w.status = resources.StatusCode(statusCode)
}

// Write composes a response and writes the response to the WaggyResponseWriter's underlying io.Writer.
// If a call to WriteHeader has not been made before calling this method, Write will call WriteHeader
// with the StatusOK (200) HTTP status code
func (w *WaggyResponseWriter) Write(body []byte) (int, error) {
	if w.status == 0 {
		w.WriteHeader(StatusOK)
	}

	if !w.Header.Has("Content-Type") {
		w.Header.Set("Content-Type", http.DetectContentType(body))
	}

	payload := w.buildResponse(body)

	return w.writer.Write(payload)
}

// Error composes a response and writes an HTTP Error Response to the WaggyResponseWriter's underlying io.Writer.
// It calls WriteHeader with the provided statusCode before composing the Error response
func (w *WaggyResponseWriter) Error(statusCode int, error string) (int, error) {
	w.WriteHeader(statusCode)

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
	w.buffer.Write([]byte(fmt.Sprintf("%s %d %s\n", os.Getenv(resources.Scheme.String()), w.status, w.status.GetStatusCodeName())))

	headerLines := make([][]byte, 0)

	for k, v := range w.Header.Loop() {
		if k == "" {
			continue
		}

		if k == resources.ContentType.String() {
			w.buffer.Write([]byte(fmt.Sprintf("%s: %s\n", k, strings.Join(v, "; "))))
		}

		w.buffer.Write([]byte(fmt.Sprintf("%s: %s\n", k, strings.Join(v, ", "))))
	}

	for _, headerLine := range headerLines {
		w.buffer.Write(headerLine)
	}

	w.buffer.Write([]byte("\n"))

	w.buffer.Write(payload)

	response := w.buffer.Bytes()

	w.buffer.Reset()

	return response
}

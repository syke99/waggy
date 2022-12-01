package v2

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestInitHandler(t *testing.T) {
	// Act
	w := InitHandler()

	// Assert
	assert.IsType(t, &WaggyHandler{}, w)
	assert.Equal(t, "", w.route)
	assert.IsType(t, []byte{}, w.defResp)
	assert.Equal(t, 0, len(w.defResp))
	assert.IsType(t, WaggyError{}, w.defErrResp)
	assert.IsType(t, map[string]http.HandlerFunc{}, w.handlerMap)
}

func TestInitHandlerWithRoute(t *testing.T) {
	// Act
	w := InitHandlerWithRoute("/test/route")

	// Assert
	assert.IsType(t, &WaggyHandler{}, w)
	assert.Equal(t, "/test/route", w.route)
	assert.IsType(t, []byte{}, w.defResp)
	assert.Equal(t, 0, len(w.defResp))
	assert.IsType(t, WaggyError{}, w.defErrResp)
	assert.IsType(t, map[string]http.HandlerFunc{}, w.handlerMap)
}

func TestWaggyHandler_WithDefaultResponse(t *testing.T) {
	// Arrange
	w := InitHandler()
	testBody := []byte("hello world")

	// Act
	w.WithDefaultResponse(testBody)

	// Assert
	assert.IsType(t, &WaggyHandler{}, w)
	assert.Equal(t, "", w.route)
	assert.IsType(t, []byte{}, w.defResp)
	assert.Equal(t, len(testBody), len(w.defResp))
	assert.Equal(t, string(testBody), string(w.defResp))
	assert.IsType(t, WaggyError{}, w.defErrResp)
	assert.IsType(t, map[string]http.HandlerFunc{}, w.handlerMap)
}

func TestWaggyHandler_WithDefaultErrorResponse(t *testing.T) {
	// Arrange
	w := InitHandler()
	testErrorType := "/test/route"
	testErrorDetail := errors.New("this is a test Waggy Error")
	testErr := WaggyError{
		Type:   testErrorType,
		Title:  "",
		Detail: testErrorDetail.Error(),
		Status: 0,
	}

	// Act
	w.WithDefaultErrorResponse(testErr, http.StatusInternalServerError)

	// Assert
	assert.IsType(t, &WaggyHandler{}, w)
	assert.Equal(t, "", w.route)
	assert.IsType(t, []byte{}, w.defResp)
	assert.Equal(t, 0, len(w.defResp))
	assert.IsType(t, WaggyError{}, w.defErrResp)
	assert.Equal(t, testErr, w.defErrResp)
	assert.Equal(t, testErrorType, w.defErrResp.Type)
	assert.Equal(t, testErrorDetail.Error(), w.defErrResp.Detail)
	assert.Equal(t, http.StatusInternalServerError, w.defErrRespCode)
	assert.IsType(t, map[string]http.HandlerFunc{}, w.handlerMap)
}

func TestWaggyHandler_MethodHandler(t *testing.T) {
	// Arrange
	helloHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello")
	}
	goodbyeHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "goodbye")
	}
	w := InitHandler()

	// Act
	w.MethodHandler(http.MethodGet, helloHandler)
	w.MethodHandler(http.MethodDelete, goodbyeHandler)

	// Assert
	assert.IsType(t, &WaggyHandler{}, w)
	assert.Equal(t, "", w.route)
	assert.IsType(t, []byte{}, w.defResp)
	assert.Equal(t, 0, len(w.defResp))
	assert.IsType(t, WaggyError{}, w.defErrResp)
	assert.IsType(t, map[string]http.HandlerFunc{}, w.handlerMap)
}

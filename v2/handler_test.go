package v2

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/syke99/waggy/v2/internal/resources"
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
	w := InitHandlerWithRoute(resources.TestRoute)

	// Assert
	assert.IsType(t, &WaggyHandler{}, w)
	assert.Equal(t, resources.TestRoute, w.route)
	assert.IsType(t, []byte{}, w.defResp)
	assert.Equal(t, 0, len(w.defResp))
	assert.IsType(t, WaggyError{}, w.defErrResp)
	assert.IsType(t, map[string]http.HandlerFunc{}, w.handlerMap)
}

func TestWaggyHandler_WithDefaultResponse(t *testing.T) {
	// Arrange
	w := InitHandler()

	// Act
	w.WithDefaultResponse([]byte(resources.HelloWorld))

	// Assert
	assert.IsType(t, &WaggyHandler{}, w)
	assert.Equal(t, "", w.route)
	assert.IsType(t, []byte{}, w.defResp)
	assert.Equal(t, len(resources.HelloWorld), len(w.defResp))
	assert.Equal(t, resources.HelloWorld, string(w.defResp))
	assert.IsType(t, WaggyError{}, w.defErrResp)
	assert.IsType(t, map[string]http.HandlerFunc{}, w.handlerMap)
}

func TestWaggyHandler_WithDefaultErrorResponse(t *testing.T) {
	// Arrange
	w := InitHandler()
	testErrorType := resources.TestRoute
	testErrorDetail := resources.TestError
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
		fmt.Fprintln(w, resources.Hello)
	}
	goodbyeHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, resources.Goodbye)
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

	for k, v := range w.handlerMap {
		switch k {
		case http.MethodGet:
			assert.Equal(t, resources.GetFunctionName(helloHandler), resources.GetFunctionName(v))
		case http.MethodDelete:
			assert.Equal(t, resources.GetFunctionName(goodbyeHandler), resources.GetFunctionName(v))
		}
	}
}

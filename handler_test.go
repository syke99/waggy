package waggy

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/syke99/waggy/internal/resources"
)

func TestAllHTTPMethods(t *testing.T) {
	// Act
	all := AllHTTPMethods()

	// Assert
	assert.Equal(t, "ALL", all)
}

func TestInitHandler(t *testing.T) {
	// Act
	w := InitHandler(nil)

	// Assert
	assert.IsType(t, &Handler{}, w)
	assert.Equal(t, "", w.route)
	assert.IsType(t, []byte{}, w.defResp)
	assert.Equal(t, 0, len(w.defResp))
	assert.IsType(t, WaggyError{}, w.defErrResp)
	assert.IsType(t, map[string]http.HandlerFunc{}, w.handlerMap)
}

func TestInitHandlerWithRoute(t *testing.T) {
	// Act
	w := InitHandlerWithRoute(resources.TestRoute, nil)

	// Assert
	assert.IsType(t, &Handler{}, w)
	assert.Equal(t, resources.TestRoute[1:], w.route)
	assert.IsType(t, []byte{}, w.defResp)
	assert.Equal(t, 0, len(w.defResp))
	assert.IsType(t, WaggyError{}, w.defErrResp)
	assert.IsType(t, map[string]http.HandlerFunc{}, w.handlerMap)
}

func TestHandler_WithDefaultResponse(t *testing.T) {
	// Arrange
	w := InitHandler(nil)

	// Act
	w.WithDefaultResponse(resources.TestContentType, []byte(resources.HelloWorld))

	// Assert
	assert.IsType(t, &Handler{}, w)
	assert.Equal(t, "", w.route)
	assert.IsType(t, []byte{}, w.defResp)
	assert.Equal(t, len(resources.HelloWorld), len(w.defResp))
	assert.Equal(t, resources.HelloWorld, string(w.defResp))
	assert.IsType(t, WaggyError{}, w.defErrResp)
	assert.IsType(t, map[string]http.HandlerFunc{}, w.handlerMap)
}

func TestHandler_WithDefaultErrorResponse(t *testing.T) {
	// Arrange
	w := InitHandler(nil)
	testErr := WaggyError{
		Type:   resources.TestRoute,
		Title:  "",
		Detail: resources.TestError.Error(),
		Status: 0,
	}

	// Act
	w.WithDefaultErrorResponse(testErr, http.StatusInternalServerError)

	// Assert
	assert.IsType(t, &Handler{}, w)
	assert.Equal(t, "", w.route)
	assert.IsType(t, []byte{}, w.defResp)
	assert.Equal(t, 0, len(w.defResp))
	assert.IsType(t, WaggyError{}, w.defErrResp)
	assert.Equal(t, testErr, w.defErrResp)
	assert.Equal(t, resources.TestRoute, w.defErrResp.Type)
	assert.Equal(t, resources.TestError.Error(), w.defErrResp.Detail)
	assert.Equal(t, http.StatusInternalServerError, w.defErrRespCode)
	assert.IsType(t, map[string]http.HandlerFunc{}, w.handlerMap)
}

func TestHandler_MethodHandler(t *testing.T) {
	// Arrange
	helloHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, resources.Hello)
	}
	goodbyeHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, resources.Goodbye)
	}
	w := InitHandler(nil)

	// Act
	w.WithMethodHandler(http.MethodGet, helloHandler)
	w.WithMethodHandler(http.MethodDelete, goodbyeHandler)

	// Assert
	assert.IsType(t, &Handler{}, w)
	assert.Equal(t, "", w.route)
	assert.IsType(t, []byte{}, w.defResp)
	assert.Equal(t, 0, len(w.defResp))
	assert.IsType(t, WaggyError{}, w.defErrResp)
	assert.IsType(t, map[string]http.HandlerFunc{}, w.handlerMap)

	for k, v := range w.handlerMap {
		switch k {
		case http.MethodGet:
			assert.Equal(t, resources.GetFunctionPtr(helloHandler), resources.GetFunctionPtr(v))
		case http.MethodDelete:
			assert.Equal(t, resources.GetFunctionPtr(goodbyeHandler), resources.GetFunctionPtr(v))
		}
	}
}

func TestHandler_RestrictMethods(t *testing.T) {
	// Arrange
	w := InitHandler(nil)

	// Act
	w = w.RestrictMethods(http.MethodGet,
		http.MethodPut,
		http.MethodPost,
		http.MethodPatch,
		http.MethodDelete,
		http.MethodConnect,
		http.MethodHead,
		http.MethodTrace,
		http.MethodOptions)

	// Assert
	for k := range resources.AllHTTPMethods() {
		_, ok := w.restrictedMethods[k]
		assert.True(t, ok)
	}
}

func TestHandler_RestrictMethods_NotHTTPMethod(t *testing.T) {
	// Arrange
	w := InitHandler(nil)
	test := "this isn't an http method"

	// Act
	w = w.RestrictMethods(test)

	_, ok := w.restrictedMethods[test]

	// Assert
	assert.Equal(t, 0, len(w.restrictedMethods))
	assert.False(t, ok)
}

func TestHandler_WithRestrictedMethodHandler(t *testing.T) {
	// Arrange
	w := InitHandler(nil)
	testHandler := func(w http.ResponseWriter, r *http.Request) {}

	// Act
	w.WithRestrictedMethodHandler(testHandler)

	// Asset
	assert.NotNil(t, w.restrictedMethodFunc)
}

func TestHandler_WithRestrictedMethodHandler_NoHandler(t *testing.T) {
	// Arrange
	w := InitHandler(nil)

	// Act
	w.WithRestrictedMethodHandler(nil)

	// Asset
	assert.Nil(t, w.restrictedMethodFunc)
}

func TestHandler_WithLogger(t *testing.T) {
	// Arrange
	testLog := resources.TestLogFile
	testLogLevel := Info
	testLogger := NewLogger(testLogLevel, testLog)

	w := InitHandler(nil)

	// Act
	w.WithLogger(testLogger, nil)

	// Assert
	assert.IsType(t, &Logger{}, w.logger)
	assert.Equal(t, Info.level(), w.logger.logLevel)
	assert.Equal(t, "", w.logger.key)
	assert.Equal(t, "", w.logger.message)
	assert.Equal(t, "", w.logger.err)
	assert.Equal(t, 0, len(w.logger.vals))
	assert.Equal(t, resources.TestLogFile, w.logger.log)
}

func TestHandler_WithLogger_ParentOverride(t *testing.T) {
	// Assert
	r := InitRouter(nil).
		WithDefaultLogger()

	testLog := resources.TestLogFile
	testLogLevel := Info
	testLogger := NewLogger(testLogLevel, testLog)

	w := InitHandler(nil)

	// Act
	w.WithLogger(testLogger, OverrideParentLogger())

	r.Handle(resources.TestRoute, w)

	// Assert
	assert.IsType(t, &Logger{}, w.logger)
	assert.Equal(t, Info.level(), w.logger.logLevel)
	assert.Equal(t, "", w.logger.key)
	assert.Equal(t, "", w.logger.message)
	assert.Equal(t, "", w.logger.err)
	assert.Equal(t, 0, len(w.logger.vals))
	assert.Equal(t, resources.TestLogFile, w.logger.log)
}

func TestHandler_WithDefaultLogger(t *testing.T) {
	// Arrange
	w := InitHandler(nil)

	// Act
	w.WithDefaultLogger()

	// Assert
	assert.IsType(t, &Logger{}, w.logger)
	assert.Equal(t, Info.level(), w.logger.logLevel)
	assert.Equal(t, "", w.logger.key)
	assert.Equal(t, "", w.logger.message)
	assert.Equal(t, "", w.logger.err)
	assert.Equal(t, 0, len(w.logger.vals))
	assert.Equal(t, os.Stderr, w.logger.log)
}

func TestHandler_Logger(t *testing.T) {
	// Arrange
	testLog := resources.TestLogFile
	testLogLevel := Info
	testLogger := NewLogger(testLogLevel, testLog)

	w := InitHandler(nil).
		WithLogger(testLogger, nil)

	// Act
	l := w.Logger()

	// Assert
	assert.IsType(t, &Logger{}, l)
	assert.Equal(t, Info.level(), l.logLevel)
	assert.Equal(t, "", l.key)
	assert.Equal(t, "", l.message)
	assert.Equal(t, "", l.err)
	assert.Equal(t, 0, len(l.vals))
	assert.Equal(t, resources.TestLogFile, l.log)
}

func TestHandler_Logger_Inherited_NoParentOverride(t *testing.T) {
	// Assert
	r := InitRouter(nil).
		WithDefaultLogger()

	testLog := resources.TestLogFile
	testLogLevel := Info
	testLogger := NewLogger(testLogLevel, testLog)

	w := InitHandler(nil)

	w.WithLogger(testLogger, nil)

	r.Handle(resources.TestRoute, w)

	// Act
	l := w.Logger()

	// Assert
	assert.IsType(t, &Logger{}, l)
	assert.Equal(t, Info.level(), l.logLevel)
	assert.Equal(t, "", l.key)
	assert.Equal(t, "", l.message)
	assert.Equal(t, "", l.err)
	assert.Equal(t, 0, len(l.vals))
	assert.Equal(t, os.Stderr, l.log)
}

func TestHandler_Logger_Inherited_ParentOverride(t *testing.T) {
	// Assert
	r := InitRouter(nil).
		WithDefaultLogger()

	testLog := resources.TestLogFile
	testLogLevel := Info
	testLogger := NewLogger(testLogLevel, testLog)

	w := InitHandler(nil)

	w.WithLogger(testLogger, OverrideParentLogger())

	r.Handle(resources.TestRoute, w)

	// Act
	l := w.Logger()

	// Assert
	assert.IsType(t, &Logger{}, l)
	assert.Equal(t, Info.level(), l.logLevel)
	assert.Equal(t, "", l.key)
	assert.Equal(t, "", l.message)
	assert.Equal(t, "", l.err)
	assert.Equal(t, 0, len(l.vals))
	assert.Equal(t, resources.TestLogFile, l.log)
}

func TestHandler_ServeHTTP_RestrictedMethod_NoHandler(t *testing.T) {
	// Arrange
	w := InitHandler(nil).
		RestrictMethods(http.MethodGet)

	r, _ := http.NewRequest(http.MethodGet, resources.TestRoute, nil)

	wr := httptest.NewRecorder()

	// Act
	w.ServeHTTP(wr, r)

	// Assert
	assert.Equal(t, http.StatusMethodNotAllowed, wr.Code)
	assert.Equal(t, resources.TestMethodNotAllowed, wr.Body.String())
}

func TestHandler_ServeHTTP_RestrictedMethod_Handler(t *testing.T) {
	// Arrange
	noRouteHandler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintln(w, "this method isn't allowed, sorry")
	}

	w := InitHandler(nil).
		RestrictMethods(http.MethodGet).
		WithRestrictedMethodHandler(noRouteHandler)

	r, _ := http.NewRequest(http.MethodGet, resources.TestRoute, nil)

	wr := httptest.NewRecorder()

	// Act
	w.ServeHTTP(wr, r)

	// Assert
	assert.Equal(t, http.StatusMethodNotAllowed, wr.Code)
	assert.Equal(t, resources.TestMethodNotAllowedHandlerResp, wr.Body.String())
}

func TestHandler_ServeHTTP_MethodGet(t *testing.T) {
	// Arrange
	helloHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, resources.Hello)
	}

	goodbyeHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, resources.Goodbye)
	}

	w := InitHandlerWithRoute(resources.TestRoute, nil)

	w.WithMethodHandler(http.MethodGet, helloHandler)
	w.WithMethodHandler(http.MethodDelete, goodbyeHandler)

	w.WithDefaultResponse(resources.TestContentType, []byte(resources.HelloWorld))

	testErr := WaggyError{
		Type:   resources.TestRoute,
		Title:  "",
		Detail: resources.TestError.Error(),
		Status: 0,
	}

	w.WithDefaultErrorResponse(testErr, http.StatusInternalServerError)

	r, _ := http.NewRequest(http.MethodGet, resources.TestRoute, nil)

	wr := httptest.NewRecorder()

	// Act
	w.ServeHTTP(wr, r)

	// Assert
	assert.IsType(t, &Handler{}, w)
	assert.Equal(t, resources.TestRoute[1:], w.route)
	assert.IsType(t, []byte{}, w.defResp)
	assert.Equal(t, len(resources.HelloWorld), len(w.defResp))
	assert.Equal(t, resources.HelloWorld, string(w.defResp))
	assert.Equal(t, testErr, w.defErrResp)
	assert.Equal(t, resources.TestRoute, w.defErrResp.Type)
	assert.Equal(t, resources.TestError.Error(), w.defErrResp.Detail)
	assert.Equal(t, http.StatusInternalServerError, w.defErrRespCode)
	assert.IsType(t, map[string]http.HandlerFunc{}, w.handlerMap)

	for k, v := range w.handlerMap {
		switch k {
		case http.MethodGet:
			assert.Equal(t, resources.GetFunctionPtr(helloHandler), resources.GetFunctionPtr(v))
		case http.MethodDelete:
			assert.Equal(t, resources.GetFunctionPtr(goodbyeHandler), resources.GetFunctionPtr(v))
		}
	}

	assert.Equal(t, fmt.Sprintf("%s\n", resources.Hello), wr.Body.String())
}

func TestHandler_ServeHTTP_MethodDelete(t *testing.T) {
	// Arrange

	helloHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, resources.Hello)
	}
	goodbyeHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, resources.Goodbye)
	}

	w := InitHandlerWithRoute(resources.TestRoute, nil)

	w.WithMethodHandler(http.MethodGet, helloHandler)
	w.WithMethodHandler(http.MethodDelete, goodbyeHandler)

	w.WithDefaultResponse(resources.TestContentType, []byte(resources.HelloWorld))

	testErr := WaggyError{
		Type:   resources.TestRoute,
		Title:  "",
		Detail: resources.TestError.Error(),
		Status: 0,
	}

	w.WithDefaultErrorResponse(testErr, http.StatusInternalServerError)

	r, _ := http.NewRequest(http.MethodDelete, resources.TestRoute, nil)

	wr := httptest.NewRecorder()

	// Act
	w.ServeHTTP(wr, r)

	// Assert
	assert.IsType(t, &Handler{}, w)
	assert.Equal(t, resources.TestRoute[1:], w.route)
	assert.IsType(t, []byte{}, w.defResp)
	assert.Equal(t, len(resources.HelloWorld), len(w.defResp))
	assert.Equal(t, resources.HelloWorld, string(w.defResp))
	assert.Equal(t, testErr, w.defErrResp)
	assert.Equal(t, resources.TestRoute, w.defErrResp.Type)
	assert.Equal(t, resources.TestError.Error(), w.defErrResp.Detail)
	assert.Equal(t, http.StatusInternalServerError, w.defErrRespCode)
	assert.IsType(t, map[string]http.HandlerFunc{}, w.handlerMap)

	for k, v := range w.handlerMap {
		switch k {
		case http.MethodGet:
			assert.Equal(t, resources.GetFunctionPtr(helloHandler), resources.GetFunctionPtr(v))
		case http.MethodDelete:
			assert.Equal(t, resources.GetFunctionPtr(goodbyeHandler), resources.GetFunctionPtr(v))
		}
	}

	assert.Equal(t, fmt.Sprintf("%s\n", resources.Goodbye), wr.Body.String())
}

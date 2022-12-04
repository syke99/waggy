package v2

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/syke99/waggy/v2/internal/resources"
	"net/http"
	"net/http/httptest"
	"os"
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
	testErr := WaggyError{
		Type:   resources.TestRoute,
		Title:  "",
		Detail: resources.TestError.Error(),
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
	assert.Equal(t, resources.TestRoute, w.defErrResp.Type)
	assert.Equal(t, resources.TestError.Error(), w.defErrResp.Detail)
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

func TestWaggyHandler_WithLogger(t *testing.T) {
	// Arrange
	testLog := resources.TestLogFile
	testLogLevel := Info
	testLogger := NewLogger(testLogLevel, testLog)

	w := InitHandler()

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

func TestWaggyHandler_WithLogger_ParentOverride(t *testing.T) {
	// Assert
	r := InitRouter().
		WithDefaultLogger()

	testLog := resources.TestLogFile
	testLogLevel := Info
	testLogger := NewLogger(testLogLevel, testLog)

	w := InitHandler()

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

func TestWaggyHandler_WithDefaultLogger(t *testing.T) {
	// Arrange
	w := InitHandler()

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

func TestWaggyHandler_Logger(t *testing.T) {
	// Arrange
	testLog := resources.TestLogFile
	testLogLevel := Info
	testLogger := NewLogger(testLogLevel, testLog)

	w := InitHandler().
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

func TestWaggyHandler_Logger_Inherited_NoParentOverride(t *testing.T) {
	// Assert
	r := InitRouter().
		WithDefaultLogger()

	testLog := resources.TestLogFile
	testLogLevel := Info
	testLogger := NewLogger(testLogLevel, testLog)

	w := InitHandler()

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

func TestWaggyHandler_Logger_Inherited_ParentOverride(t *testing.T) {
	// Assert
	r := InitRouter().
		WithDefaultLogger()

	testLog := resources.TestLogFile
	testLogLevel := Info
	testLogger := NewLogger(testLogLevel, testLog)

	w := InitHandler()

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

func TestWaggyHandler_FileServer(t *testing.T) {
	// Arrange
	testPath := resources.TestFilePath
	wr := httptest.NewRecorder()

	w := InitHandler().
		WithDefaultLogger()

	// Act
	w.ServeFile(wr, testPath)

	// Assert
	assert.Equal(t, http.StatusOK, wr.Code)
}

func TestWaggyHandler_FileServer_NoPath(t *testing.T) {
	// Arrange
	wr := httptest.NewRecorder()

	w := InitHandler().
		WithDefaultLogger()

	// Act
	w.ServeFile(wr, "")

	// Assert
	assert.Equal(t, http.StatusInternalServerError, wr.Code)
	assert.Equal(t, "Error serving file", wr.Body.String())
}

func TestWaggyHandler_ServeHTTP_MethodGet(t *testing.T) {
	// Arrange
	os.Setenv(resources.XMatchedRoute.String(), resources.TestRoute)
	os.Setenv(resources.RequestMethod.String(), http.MethodGet)

	helloHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, resources.Hello)
	}
	goodbyeHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, resources.Goodbye)
	}

	w := InitHandlerWithRoute(resources.TestRoute)

	w.MethodHandler(http.MethodGet, helloHandler)
	w.MethodHandler(http.MethodDelete, goodbyeHandler)

	w.WithDefaultResponse([]byte(resources.HelloWorld))

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
	assert.IsType(t, &WaggyHandler{}, w)
	assert.Equal(t, resources.TestRoute, w.route)
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
			assert.Equal(t, resources.GetFunctionName(helloHandler), resources.GetFunctionName(v))
		case http.MethodDelete:
			assert.Equal(t, resources.GetFunctionName(goodbyeHandler), resources.GetFunctionName(v))
		}
	}

	assert.Equal(t, fmt.Sprintf("%s\n", resources.Hello), wr.Body.String())
}

func TestWaggyHandler_ServeHTTP_MethodDelete(t *testing.T) {
	// Arrange
	os.Setenv(resources.XMatchedRoute.String(), resources.TestRoute)
	os.Setenv(resources.RequestMethod.String(), http.MethodDelete)

	helloHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, resources.Hello)
	}
	goodbyeHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, resources.Goodbye)
	}

	w := InitHandlerWithRoute(resources.TestRoute)

	w.MethodHandler(http.MethodGet, helloHandler)
	w.MethodHandler(http.MethodDelete, goodbyeHandler)

	w.WithDefaultResponse([]byte(resources.HelloWorld))

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
	assert.IsType(t, &WaggyHandler{}, w)
	assert.Equal(t, resources.TestRoute, w.route)
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
			assert.Equal(t, resources.GetFunctionName(helloHandler), resources.GetFunctionName(v))
		case http.MethodDelete:
			assert.Equal(t, resources.GetFunctionName(goodbyeHandler), resources.GetFunctionName(v))
		}
	}

	assert.Equal(t, fmt.Sprintf("%s\n", resources.Goodbye), wr.Body.String())
}

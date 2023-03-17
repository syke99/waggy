package waggy

import (
	"fmt"
	"github.com/syke99/waggy/internal/json"
	"github.com/syke99/waggy/middleware"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/syke99/waggy/internal/resources"
)

func TestLog(t *testing.T) {
	// Arrange
	logger := Logger{
		logLevel: "",
		key:      "",
		message:  "",
		err:      "",
		vals:     make(map[string]interface{}),
		log:      os.Stdout,
	}

	greetingHandler := func(w http.ResponseWriter, r *http.Request) {
		l := Log(r)

		// Assert
		assert.NotNil(t, l)
	}

	handler := NewHandlerWithRoute(resources.TestRoutePathParams, nil).
		WithLogger(&logger, nil)

	handler.WithMethodHandler(http.MethodGet, greetingHandler)

	r, _ := http.NewRequest(http.MethodGet, resources.TestRoutePathParamGoodbye, nil)

	wr := httptest.NewRecorder()

	handler.ServeHTTP(wr, r)
}

var testMiddleWareTest middleware.MiddleWare = func(handler http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("test", "hello world")

		handler.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func TestServeThroughMiddleWare(t *testing.T) {
	// Arrange
	testHandler := func(w http.ResponseWriter, r *http.Request) {
		// Assert
		assert.Equal(t, "hello world", w.Header().Get("test"))
	}

	handler := InitHandlerWithRoute(resources.TestRoutePathParams, nil)

	handler.WithMethodHandler(http.MethodGet, testHandler).
		Use(testMiddleWareTest)

	r, _ := http.NewRequest(http.MethodGet, resources.TestRoutePathParamGoodbye, nil)

	wr := httptest.NewRecorder()

	handler.ServeHTTP(wr, r)
}

func TestLog_Error(t *testing.T) {
	// Arrange
	greetingHandler := func(w http.ResponseWriter, r *http.Request) {
		l := Log(r)

		// Assert
		assert.Nil(t, l)
	}

	handler := NewHandlerWithRoute(resources.TestRoutePathParams, nil).
		WithLogger(nil, nil)

	handler.WithMethodHandler(http.MethodGet, greetingHandler)

	r, _ := http.NewRequest(http.MethodGet, resources.TestRoutePathParamGoodbye, nil)

	wr := httptest.NewRecorder()

	handler.ServeHTTP(wr, r)
}

func TestVars_Hello(t *testing.T) {
	// Arrange

	greetingHandler := func(w http.ResponseWriter, r *http.Request) {
		params := Vars(r)

		switch params["param"] {
		case resources.Hello:
			fmt.Fprintln(w, resources.Hello)
		case resources.Goodbye:
			fmt.Fprintln(w, resources.Goodbye)
		case "":
			fmt.Fprintln(w, resources.WhereAmI)
		}
	}

	handler := NewHandlerWithRoute(resources.TestRoutePathParams, nil)

	handler.WithMethodHandler(http.MethodGet, greetingHandler)

	r, _ := http.NewRequest(http.MethodGet, resources.TestRoutePathParamHello, nil)

	wr := httptest.NewRecorder()

	handler.ServeHTTP(wr, r)

	assert.Equal(t, fmt.Sprintf("%s\n", resources.Hello), wr.Body.String())
	assert.Equal(t, http.StatusOK, wr.Code)
}

func TestVars_Goodbye(t *testing.T) {
	// Arrange
	greetingHandler := func(w http.ResponseWriter, r *http.Request) {
		params := Vars(r)

		switch params["param"] {
		case resources.Hello:
			fmt.Fprintln(w, resources.Hello)
		case resources.Goodbye:
			fmt.Fprintln(w, resources.Goodbye)
		case "":
			fmt.Fprintln(w, resources.WhereAmI)
		}
	}

	handler := NewHandlerWithRoute(resources.TestRoutePathParams, nil)

	handler.WithMethodHandler(http.MethodGet, greetingHandler)

	r, _ := http.NewRequest(http.MethodGet, resources.TestRoutePathParamGoodbye, nil)

	wr := httptest.NewRecorder()

	// Act
	handler.ServeHTTP(wr, r)

	// Assert
	assert.Equal(t, fmt.Sprintf("%s\n", resources.Goodbye), wr.Body.String())
}

func TestVars_NoPathParams(t *testing.T) {
	// Arrange

	greetingHandler := func(w http.ResponseWriter, r *http.Request) {
		params := Vars(r)

		switch params["param"] {
		case resources.Hello:
			fmt.Fprintln(w, resources.Hello)
		case resources.Goodbye:
			fmt.Fprintln(w, resources.Goodbye)
		case "":
			fmt.Fprintln(w, resources.WhereAmI)
		}
	}

	handler := NewHandlerWithRoute(resources.TestRoute, nil)

	handler.WithMethodHandler(http.MethodGet, greetingHandler)

	r, _ := http.NewRequest(http.MethodGet, resources.TestRoute, nil)

	wr := httptest.NewRecorder()

	// Act
	handler.ServeHTTP(wr, r)

	// Assert
	assert.Equal(t, fmt.Sprintf("%s\n", resources.WhereAmI), wr.Body.String())
}

func TestWriteDefaultResponse(t *testing.T) {
	// Arrange

	defRespHandler := func(w http.ResponseWriter, r *http.Request) {
		WriteDefaultResponse(w, r)
	}

	handler := NewHandlerWithRoute(resources.TestRoute, nil).
		WithMethodHandler(http.MethodGet, defRespHandler).
		WithDefaultResponse(resources.TestContentType, []byte(resources.Hello))

	r, _ := http.NewRequest(http.MethodGet, resources.TestRoutePathParamHello, nil)

	wr := httptest.NewRecorder()

	// Act
	handler.ServeHTTP(wr, r)

	// Assault
	assert.Equal(t, fmt.Sprintf("%s\n", resources.Hello), wr.Body.String())
}

func TestWriteDefaultErrorResponse(t *testing.T) {
	// Arrange
	testErr := WaggyError{
		Type:   resources.TestRoute,
		Title:  "",
		Detail: resources.TestError.Error(),
		Status: 0,
	}
	expectedResult := fmt.Sprintf("%s\n", json.BuildJSONStringFromWaggyError(testErr.Type, "", testErr.Detail, 0, "", ""))

	defRespHandler := func(w http.ResponseWriter, r *http.Request) {
		WriteDefaultErrorResponse(w, r)
	}

	handler := NewHandlerWithRoute(resources.TestRoute, nil).
		WithMethodHandler(http.MethodGet, defRespHandler).
		WithDefaultErrorResponse(testErr, http.StatusInternalServerError)

	r, _ := http.NewRequest(http.MethodGet, resources.TestRoutePathParamHello, nil)

	wr := httptest.NewRecorder()

	// Act
	handler.ServeHTTP(wr, r)

	// Assert
	assert.Equal(t, expectedResult, wr.Body.String())
}

func TestServeFile(t *testing.T) {
	// Arrange
	wr := httptest.NewRecorder()

	// Act
	ServeFile(wr, resources.TestContentType, resources.TestFilePath)

	// Assert
	assert.Equal(t, http.StatusOK, wr.Code)
	assert.Equal(t, "application/json", wr.Header().Get("content-type"))
}

func TestServeFile_NoContentType(t *testing.T) {
	// Arrange
	wr := httptest.NewRecorder()

	// Act
	ServeFile(wr, "", resources.TestFilePath)

	// Assert
	assert.Equal(t, http.StatusOK, wr.Code)
	assert.Equal(t, "application/octet-stream", wr.Header().Get("content-type"))
}

func TestServeFile_NoPathToFile(t *testing.T) {
	// Arrange
	wr := httptest.NewRecorder()

	// Act
	ServeFile(wr, resources.TestContentType, "")

	// Assert
	assert.Equal(t, http.StatusNotFound, wr.Code)
	assert.Equal(t, "application/problem+json", wr.Header().Get("content-type"))
	assert.Equal(t, resources.TestErrorResponse, wr.Body.String())
}

func TestServe_Router(t *testing.T) {
	// Arrange
	w := NewRouter(nil)

	// Act
	err := Serve(w)

	// Assert
	assert.Error(t, err)
}

func TestServe_Handler(t *testing.T) {
	// Arrange
	w := NewHandler(nil)

	// Act
	err := Serve(w)

	// Assert
	assert.Error(t, err)
}

func TestEmptyRoute(t *testing.T) {
	handler := NewHandlerWithRoute("/", nil)
	handler.WithMethodHandler(http.MethodGet, func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "Don't panic! 🐿️")
	})
	r, _ := http.NewRequest(http.MethodGet, resources.TestRoutePathParamGoodbye, nil)
	wr := httptest.NewRecorder()
	handler.ServeHTTP(wr, r)
}

func TestListenAndServe_Router(t *testing.T) { // Act
	err := ListenAndServe[*Router](resources.TestListenAndServeAddr, nil)

	// Assert
	assert.Error(t, err)
}

func TestListenAndServe_Handler(t *testing.T) {
	// Act
	err := ListenAndServe[*Handler](resources.TestListenAndServeAddr, nil)

	// Assert
	assert.Error(t, err)
}

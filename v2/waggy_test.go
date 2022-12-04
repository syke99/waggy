package v2

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/syke99/waggy/v2/internal/resources"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestVars_Hello(t *testing.T) {
	// Arrange
	os.Setenv(resources.XMatchedRoute.String(), resources.TestRoutePathParamHello)
	os.Setenv(resources.RequestMethod.String(), http.MethodGet)

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

	handler := InitHandlerWithRoute(resources.TestRoutePathParams)

	handler.MethodHandler(http.MethodGet, greetingHandler)

	r, _ := http.NewRequest(http.MethodGet, resources.TestRoutePathParamHello, nil)

	wr := httptest.NewRecorder()

	handler.ServeHTTP(wr, r)

	assert.Equal(t, fmt.Sprintf("%s\n", resources.Hello), wr.Body.String())
	assert.Equal(t, http.StatusOK, wr.Code)
}

func TestVars_Goodbye(t *testing.T) {
	// Arrange
	os.Setenv(resources.XMatchedRoute.String(), resources.TestRoutePathParamGoodbye)
	os.Setenv(resources.RequestMethod.String(), http.MethodGet)

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

	handler := InitHandlerWithRoute(resources.TestRoutePathParams)

	handler.MethodHandler(http.MethodGet, greetingHandler)

	r, _ := http.NewRequest(http.MethodGet, resources.TestRoutePathParamGoodbye, nil)

	wr := httptest.NewRecorder()

	// Act
	handler.ServeHTTP(wr, r)

	// Assert
	assert.Equal(t, fmt.Sprintf("%s\n", resources.Goodbye), wr.Body.String())
}

func TestVars_NoPathParams(t *testing.T) {
	// Arrange
	os.Setenv(resources.XMatchedRoute.String(), resources.TestRoute)
	os.Setenv(resources.RequestMethod.String(), http.MethodGet)

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

	handler := InitHandlerWithRoute(resources.TestRoute)

	handler.MethodHandler(http.MethodGet, greetingHandler)

	r, _ := http.NewRequest(http.MethodGet, resources.TestRoute, nil)

	wr := httptest.NewRecorder()

	// Act
	handler.ServeHTTP(wr, r)

	// Assert
	assert.Equal(t, fmt.Sprintf("%s\n", resources.WhereAmI), wr.Body.String())
}

func TestWriteDefaultResponse(t *testing.T) {
	// Arrange
	os.Setenv(resources.XMatchedRoute.String(), resources.TestRoutePathParamHello)
	os.Setenv(resources.RequestMethod.String(), http.MethodGet)

	defRespHandler := func(w http.ResponseWriter, r *http.Request) {
		WriteDefaultResponse(w, r)
	}

	handler := InitHandlerWithRoute(resources.TestRoute).
		MethodHandler(http.MethodGet, defRespHandler).
		WithDefaultResponse([]byte(resources.Hello))

	r, _ := http.NewRequest(http.MethodGet, resources.TestRoutePathParamHello, nil)

	wr := httptest.NewRecorder()

	// Act
	handler.ServeHTTP(wr, r)

	// Assault
	assert.Equal(t, fmt.Sprintf("%s\n", resources.Hello), wr.Body.String())
}

func TestWriteDefaultErrorResponse(t *testing.T) {
	// Arrange
	os.Setenv(resources.XMatchedRoute.String(), resources.TestRoutePathParamHello)
	os.Setenv(resources.RequestMethod.String(), http.MethodGet)

	testErr := WaggyError{
		Type:   resources.TestRoute,
		Title:  "",
		Detail: resources.TestError.Error(),
		Status: 0,
	}

	defRespHandler := func(w http.ResponseWriter, r *http.Request) {
		WriteDefaultErrorResponse(w, r)
	}

	handler := InitHandlerWithRoute(resources.TestRoute).
		MethodHandler(http.MethodGet, defRespHandler).
		WithDefaultErrorResponse(testErr, http.StatusInternalServerError)

	r, _ := http.NewRequest(http.MethodGet, resources.TestRoutePathParamHello, nil)

	wr := httptest.NewRecorder()

	errBytes, _ := json.Marshal(testErr)

	// Act
	handler.ServeHTTP(wr, r)

	// Assert
	assert.Equal(t, fmt.Sprintf("%s\n", string(errBytes)), wr.Body.String())
}

func TestServe_Router(t *testing.T) {
	// Arrange
	w := InitRouter()

	// Act
	err := Serve(w)

	// Assert
	assert.Error(t, err)
}

func TestServe_Handler(t *testing.T) {
	// Arrange
	w := InitHandler()

	// Act
	err := Serve(w)

	// Assert
	assert.Error(t, err)
}

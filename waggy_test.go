package waggy

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/syke99/waggy/internal/resources"
)

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

	handler.buildErrorJSON()

	// Act
	handler.ServeHTTP(wr, r)

	// Assert
	assert.Equal(t, fmt.Sprintf("%s\n", handler.buildErrorJSON()), wr.Body.String())
}

func TestServe_Router(t *testing.T) {
	// Arrange
	w := InitRouter(nil)

	// Act
	err := Serve(w)

	// Assert
	assert.Error(t, err)
}

func TestServe_Handler(t *testing.T) {
	// Arrange
	w := InitHandler(nil)

	// Act
	err := Serve(w)

	// Assert
	assert.Error(t, err)
}

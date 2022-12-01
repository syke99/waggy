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
		}
	}

	handler := InitHandlerWithRoute(resources.TestRoutePathParams)

	handler.MethodHandler(http.MethodGet, greetingHandler)

	r, _ := http.NewRequest(http.MethodGet, resources.TestRoutePathParamHello, nil)

	wr := httptest.NewRecorder()

	handler.ServeHTTP(wr, r)

	assert.Equal(t, fmt.Sprintf("%s\n", resources.Hello), wr.Body.String())
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
		}
	}

	handler := InitHandlerWithRoute(resources.TestRoutePathParams)

	handler.MethodHandler(http.MethodGet, greetingHandler)

	r, _ := http.NewRequest(http.MethodGet, resources.TestRoutePathParamGoodbye, nil)

	wr := httptest.NewRecorder()

	handler.ServeHTTP(wr, r)

	assert.Equal(t, fmt.Sprintf("%s\n", resources.Goodbye), wr.Body.String())
}

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

func TestNewRouter(t *testing.T) {
	// Act
	w := NewRouter(nil)

	// Assert
	assert.IsType(t, &Router{}, w)
	assert.IsType(t, map[string]*Handler{}, w.router)
	assert.Equal(t, 0, len(w.router))
}

func TestNewRouter_Flg_Parsable(t *testing.T) {
	// Act
	var flg FullServer = "1"
	w := NewRouter(&flg)

	// Assert
	assert.IsType(t, &Router{}, w)
	assert.IsType(t, map[string]*Handler{}, w.router)
	assert.Equal(t, 0, len(w.router))
	assert.True(t, w.FullServer)
}

func TestNewRouter_Flg_NotParsable(t *testing.T) {
	// Act
	var flg FullServer = "adsf"
	w := NewRouter(&flg)

	// Assert
	assert.IsType(t, &Router{}, w)
	assert.IsType(t, map[string]*Handler{}, w.router)
	assert.Equal(t, 0, len(w.router))
	assert.False(t, w.FullServer)
}

func TestRouter_Handle(t *testing.T) {
	// Arrange
	w := NewRouter(nil)

	helloHandler := NewHandler(nil)
	goodbyeHandler := NewHandler(nil)

	// Act
	w.Handle("/hello", helloHandler)
	w.Handle("/goodbye", goodbyeHandler)

	// Assert
	assert.IsType(t, &Handler{}, w.router["/hello"])
	assert.IsType(t, &Handler{}, w.router["/goodbye"])
}

func TestRouter_WithLogger(t *testing.T) {
	// Arrange
	w := NewRouter(nil)

	testLog := resources.TestLogFile
	testLogLevel := Info
	l := NewLogger(testLogLevel, testLog)

	// Act
	w.WithLogger(l)

	// Assert
	assert.IsType(t, &Logger{}, w.logger)
	assert.Equal(t, Info.level(), w.logger.logLevel)
	assert.Equal(t, "", w.logger.key)
	assert.Equal(t, "", w.logger.message)
	assert.Equal(t, "", w.logger.err)
	assert.Equal(t, 0, len(w.logger.vals))
	assert.Equal(t, resources.TestLogFile, w.logger.log)
}

func TestRouter_WithDefaultLogger(t *testing.T) {
	// Arrange
	w := NewRouter(nil)

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

func TestRouter_Logger(t *testing.T) {
	// Arrange
	testLog := resources.TestLogFile
	testLogLevel := Info
	testLogger := NewLogger(testLogLevel, testLog)

	w := NewRouter(nil).
		WithLogger(testLogger)

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

func TestRouter_Logger_Default(t *testing.T) {
	// Arrange
	w := NewRouter(nil).
		WithDefaultLogger()

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

func TestRouter_ServeHTTP_NoBaseRoute(t *testing.T) {
	// Arrange
	goodbyeHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, resources.Goodbye)
	}

	wr := NewRouter(nil).WithNoRouteHandler(goodbyeHandler)

	helloHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, resources.Hello)
	}

	wh := NewHandler(nil).
		WithMethodHandler(http.MethodGet, helloHandler)

	wr.Handle(resources.TestRoute, wh)

	r, _ := http.NewRequest(http.MethodGet, "/", nil)

	rr := httptest.NewRecorder()

	// Act
	wr.ServeHTTP(rr, r)

	// Assert
	assert.IsType(t, &Router{}, wr)
	assert.Equal(t, http.StatusMethodNotAllowed, rr.Code)
	assert.Equal(t, fmt.Sprintf("%s\n", resources.Goodbye), rr.Body.String())
}

func TestRouter_ServeHTTP_BaseRoute(t *testing.T) {
	// Arrange
	goodbyeHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, resources.Goodbye)
	}

	wr := NewRouter(nil)

	helloHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, resources.Hello)
	}

	wh := NewHandler(nil).
		WithMethodHandler(http.MethodGet, helloHandler)

	wh2 := NewHandler(nil).WithMethodHandler(http.MethodGet, goodbyeHandler)

	wr.Handle(resources.TestRoute, wh).Handle("/", wh2)

	r, _ := http.NewRequest(http.MethodGet, "/", nil)

	rr := httptest.NewRecorder()

	// Act
	wr.ServeHTTP(rr, r)

	// Assert
	assert.IsType(t, &Router{}, wr)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, fmt.Sprintf("%s\n", resources.Goodbye), rr.Body.String())
}

func TestRouter_ServeHTTP_MethodGet(t *testing.T) {
	// Arrange
	wr := NewRouter(nil)

	helloHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, resources.Hello)
	}

	wh := NewHandler(nil).
		WithMethodHandler(http.MethodGet, helloHandler)

	wr.Handle(resources.TestRoute, wh)

	r, _ := http.NewRequest(http.MethodGet, resources.TestRoute, nil)

	rr := httptest.NewRecorder()

	// Act
	wr.ServeHTTP(rr, r)

	// Assert
	assert.IsType(t, &Router{}, wr)
	assert.IsType(t, &Handler{}, wr.router[resources.TestRoute])
	assert.Equal(t, fmt.Sprintf("%s\n", resources.Hello), rr.Body.String())
}

func TestRouter_ServeHTTP_MethodDelete(t *testing.T) {
	// Arrange
	wr := NewRouter(nil)

	goodbyeHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, resources.Goodbye)
	}

	wh := NewHandler(nil).
		WithMethodHandler(http.MethodDelete, goodbyeHandler)

	wr.Handle(resources.TestRoute, wh)

	r, _ := http.NewRequest(http.MethodDelete, resources.TestRoute, nil)

	rr := httptest.NewRecorder()

	// Act
	wr.ServeHTTP(rr, r)

	// Assert
	assert.IsType(t, &Router{}, wr)
	assert.IsType(t, &Handler{}, wr.router[resources.TestRoute])
	assert.Equal(t, fmt.Sprintf("%s\n", resources.Goodbye), rr.Body.String())
}

func TestRouter_Walk(t *testing.T) {
	// Arrange
	wr := NewRouter(nil)

	helloHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, resources.Hello)
	}

	goodbyeHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, resources.Goodbye)
	}

	wh := NewHandler(nil).
		WithMethodHandler(http.MethodDelete, goodbyeHandler)

	wr.Handle(resources.TestRoute, wh)

	wh2 := NewHandler(nil).
		WithMethodHandler(http.MethodGet, helloHandler)

	wr.Handle(resources.TestRouteTwo, wh2)

	methods := []string{http.MethodDelete, http.MethodGet}
	routes := []string{resources.TestRoute, resources.TestRouteTwo}

	// Act
	_ = wr.Walk(func(method string, route string) error {
		index := func(m string, methods []string) int {
			for i, mthd := range methods {
				if mthd == m {
					return i
				}
			}
			return 0
		}

		i := index(method, methods)

		// Assert
		assert.Equal(t, route, routes[i])

		return nil
	})
}

func TestRouter_Walk_Error(t *testing.T) {
	// Arrange
	wr := NewRouter(nil)

	helloHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, resources.Hello)
	}

	goodbyeHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, resources.Goodbye)
	}

	wh := NewHandler(nil).
		WithMethodHandler(http.MethodDelete, goodbyeHandler)

	wr.Handle(resources.TestRoute, wh)

	wh2 := NewHandler(nil).
		WithMethodHandler(http.MethodGet, helloHandler)

	wr.Handle(resources.TestRouteThree, wh2)

	// Act
	err := wr.Walk(func(method string, route string) error {
		index := func(m string, methods []string) int {
			for i, mthd := range methods {
				if mthd == m {
					return i
				}
			}
			return 0
		}

		if index(method, resources.TestMethods) == 0 {
			return resources.TestError
		}

		return nil
	})

	// Assert
	assert.Error(t, err)
	assert.Equal(t, resources.TestError, err)
}

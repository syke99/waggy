package waggy

import (
	"errors"
	"fmt"
	"github.com/syke99/waggy/middleware"
	"io"
	"net/http"
	"net/http/cgi"
	"os"

	"github.com/syke99/waggy/internal/resources"
)

type FullServer string

// WaggyEntryPoint is used as a type constraint whenever calling
// Serve so that only a *Router or *Handler can
// be used and not a bare http.Handler
type WaggyEntryPoint interface {
	*Router | *Handler
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

// WriteDefaultResponse returns the result (number of bytes written
// and a nil value, or the error of that write) of writing the set
// default response inside the handler it is being used inside of.
// If no default response has been set, this function will return
// an error.
func WriteDefaultResponse(w http.ResponseWriter, r *http.Request) {
	rv := r.Context().Value(resources.DefResp)
	if rv == nil {
		fmt.Fprintln(w, resources.NoDefaultResponse.Error())
	}

	fn := rv.(func(wr http.ResponseWriter))

	fn(w)
}

// WriteDefaultErrorResponse returns the result of writing the set
// default error response inside the handler it is being used inside of.
// If no default error response has been set, this function will return
// an error.
func WriteDefaultErrorResponse(w http.ResponseWriter, r *http.Request) {
	rv := r.Context().Value(resources.DefErr)
	if rv == nil {
		fmt.Fprintln(w, resources.NoDefaultErrorResponse.Error())
	}

	fn := rv.(func(wr http.ResponseWriter))

	fn(w)
}

// Log returns the *Logger for the current request handler
// to log a message to the set *os.File (defautls to os.Stderr)
// if one exists, otherwise nil
func Log(r *http.Request) *Logger {
	if rv := r.Context().Value(resources.Logger); rv != nil {
		return rv.(*Logger)
	}
	return nil
}

// Vars returns the values matching any path parameters if they exist,
// otherwise returns nil
func Vars(r *http.Request) map[string]string {
	if rv := r.Context().Value(resources.PathParams); rv != nil {
		return rv.(map[string]string)
	}
	return nil
}

// Serve wraps a call to cgi.serve and also uses a type constraint of
// WaggyEntryPoint so that only a *Router or *Handler can be
// used in the call to Serve and not accidentally allow calling
// a bare http.Handler
func Serve[W WaggyEntryPoint](entryPoint W) error {
	return cgi.Serve(entryPoint)
}

// ListenAndServe wraps a call to http.ListenAndServe and also uses
// a type constraint of WaggyEntryPoint so that only a *waggy.Router or
// *waggy.Handler can be used
func ListenAndServe[W WaggyEntryPoint](addr string, entryPoint W) error {
	if entryPoint == nil {
		return resources.NoWaggyEntryPointProvided
	}
	return http.ListenAndServe(addr, entryPoint)
}

// ServeFile is a convenience function for serving the file at the given filePath to the given
// http.ResponseWriter (w). If Waggy cannot find a file at the given path (if it doesn't exist
// or the volume was incorrectly mounted), this function will return a status 404. If any other
// error occurs, this function will return a 500. If no contentType is given, this function will
// set the Content-Type header to "application/octet-stream"
func ServeFile(w http.ResponseWriter, contentType string, filePath string) {
	var err error

	errMsg := WaggyError{
		Title:  "",
		Status: 0,
	}

	if filePath == "" {
		err = errors.New("no path to file provided")
		errMsg.Title = "Resource Not Found"
		errMsg.Status = http.StatusNotFound
	}

	file := new(os.File)
	if err == nil {
		file, err = os.Open(filePath)
	}

	if err == nil {
		if contentType == "" {
			contentType = "application/octet-stream"
		}
		w.Header().Set("content-type", contentType)
		_, err = io.Copy(w, file)
	}

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		if errMsg.Status == 0 {
			errMsg.Status = http.StatusInternalServerError
			errMsg.Title = "Internal Server Error"
			w.WriteHeader(http.StatusInternalServerError)
		}

		errJSON := fmt.Sprintf("{ \"title\": \"%[1]s\", \"detail\": \"%[2]s\", \"status\": \"%[3]d\" }", errMsg.Title, err.Error(), errMsg.Status)

		w.Header().Set("content-type", "application/problem+json")
		fmt.Fprint(w, errJSON)
	}
}

func serveThroughMiddleWare(middle []middleware.MiddleWare, handler http.HandlerFunc) http.HandlerFunc {
	for _, mw := range middle {
		handler = mw(handler).ServeHTTP
	}

	return handler
}

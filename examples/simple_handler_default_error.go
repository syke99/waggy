package examples

import (
	"fmt"
	wagi "github.com/syke99/waggy/v2"
	"net/http"
)

func defErrorHandler(w http.ResponseWriter, r *http.Request) {
	params := wagi.Vars(r)

	greetingType := params["type"]

	switch greetingType {
	case "hello":
		fmt.Fprintln(w, "Hello World!!")
	case "goodbye":
		fmt.Fprintln(w, "Goodbye for now!!!")
	case "whatup":
		wagi.WriteDefaultErrorResponse(w, r)
	case "":
		wagi.WriteDefaultErrorResponse(w, r)
	}
}

func ExampleDefaultError() {
	defaultError := wagi.WaggyError{
		Type:   "/greeting",
		Detail: "no type parameter provided",
		Status: http.StatusBadRequest,
	}

	greetingHandler := wagi.InitHandlerWithRoute("/greeting/{type}").
		MethodHandler(http.MethodGet, defErrorHandler).
		WithDefaultResponse([]byte("So what's good?")).
		WithDefaultErrorResponse(defaultError, http.StatusBadRequest)

	_ = wagi.Serve(greetingHandler)
}

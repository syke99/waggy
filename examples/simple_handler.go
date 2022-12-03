package examples

import (
	"fmt"
	wagi "github.com/syke99/waggy/v2"
	"net/http"
)

func handlerHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello")
}

func handlerGoodbye(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Goodbye")
}

func ExampleHandler() {
	greetingHandler := wagi.InitHandlerWithRoute("/greeting").
		MethodHandler(http.MethodGet, handlerHello).
		MethodHandler(http.MethodDelete, handlerGoodbye)

	_ = wagi.Serve(greetingHandler)
}

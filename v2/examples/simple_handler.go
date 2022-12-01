package examples

import (
	"fmt"
	"net/http"
	"net/http/cgi"

	wagi "github.com/syke99/waggy/v2"
)

func HandlerHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello")
}

func HandlerGoodbye(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Goodbye")
}

func main() {
	greetingHandler := wagi.InitHandlerWithRoute("/greeting").
		MethodHandler(http.MethodGet, HandlerHello).
		MethodHandler(http.MethodDelete, HandlerGoodbye)

	_ = cgi.Serve(greetingHandler)
}

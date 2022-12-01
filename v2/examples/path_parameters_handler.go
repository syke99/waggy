package examples

import (
	"fmt"
	"net/http"
	"net/http/cgi"

	wagi "github.com/syke99/waggy/v2"
)

func GreetingHandler(w http.ResponseWriter, r *http.Request) {
	params := wagi.Vars(r)

	greetingType := params["type"]

	switch greetingType {
	case "hello":
		fmt.Fprintln(w, "Hello World!!")
	case "goodbye":
		fmt.Fprintln(w, "Goodbye for now!!!")
	}
}

func main() {
	greetingHandler := wagi.InitHandlerWithRoute("/greeting/{type}").
		MethodHandler(http.MethodGet, GreetingHandler)

	_ = cgi.Serve(greetingHandler)
}

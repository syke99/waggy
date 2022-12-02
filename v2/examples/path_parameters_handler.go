package examples

import (
	"fmt"
	wagi "github.com/syke99/waggy/v2"
	"net/http"
)

func HandlerPathParamsHandler(w http.ResponseWriter, r *http.Request) {
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
		MethodHandler(http.MethodGet, HandlerPathParamsHandler)

	_ = wagi.Serve(greetingHandler)
}

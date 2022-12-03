package examples

import (
	"fmt"
	wagi "github.com/syke99/waggy/v2"
	"net/http"
)

func handlerPathParamsHandler(w http.ResponseWriter, r *http.Request) {
	params := wagi.Vars(r)

	greetingType := params["type"]

	switch greetingType {
	case "hello":
		fmt.Fprintln(w, "Hello World!!")
	case "goodbye":
		fmt.Fprintln(w, "Goodbye for now!!!")
	}
}

func ExampleHandlerPathParams() {
	greetingHandler := wagi.InitHandlerWithRoute("/greeting/{type}").
		MethodHandler(http.MethodGet, handlerPathParamsHandler)

	_ = wagi.Serve(greetingHandler)
}

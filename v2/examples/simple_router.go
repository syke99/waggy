package examples

import (
	"fmt"
	wagi "github.com/syke99/waggy/v2"
	"net/http"
)

func routerHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello")
}

func routerGoodbye(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Goodbye")
}

func ExampleRouter() {
	greetingHandler := wagi.InitHandler().
		MethodHandler(http.MethodGet, routerHello).
		MethodHandler(http.MethodDelete, routerGoodbye)

	router := wagi.InitRouter()

	router.Handle("/greeting", greetingHandler)

	_ = wagi.Serve(router)
}

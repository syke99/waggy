package examples

import (
	"fmt"
	wagi "github.com/syke99/waggy/v2"
	"net/http"
)

func RouterHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello")
}

func RouterGoodbye(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Goodbye")
}

func main() {
	greetingHandler := wagi.InitHandler().
		MethodHandler(http.MethodGet, RouterHello).
		MethodHandler(http.MethodDelete, RouterGoodbye)

	router := wagi.InitRouter()

	router.Handle("/greeting", greetingHandler)

	_ = wagi.Serve(router)
}

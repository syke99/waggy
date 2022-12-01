package examples

import (
	"fmt"
	"net/http"
	"net/http/cgi"

	wagi "github.com/syke99/waggy/v2"
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

	_ = cgi.Serve(router)
}

package examples

import (
	"fmt"
	"github.com/syke99/waggy/v2"
	"net/http"
	"net/http/cgi"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello")
}

func Goodbye(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Goodbye")
}

func main() {
	greetingHandler := v2.InitHandler().
		MethodHandler(http.MethodGet, Hello).
		MethodHandler(http.MethodDelete, Goodbye)

	router := v2.InitRouter()

	router.Handle("/greeting", greetingHandler)

	_ = cgi.Serve(router)
}

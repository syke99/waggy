package examples

import (
	"fmt"
	"github.com/syke99/waggy/v2/waggy"
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
	greetingHandler := waggy.InitHandler().
		MethodHandler(http.MethodGet, Hello).
		MethodHandler(http.MethodDelete, Goodbye)

	router := waggy.InitRouter()

	router.Handle("/greeting", greetingHandler)

	_ = cgi.Serve(router)
}

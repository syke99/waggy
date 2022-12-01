package examples

import (
	"fmt"
	"net/http"
	"net/http/cgi"

	wagi "github.com/syke99/waggy/v2"
)

func DefRespHandler(w http.ResponseWriter, r *http.Request) {
	params := wagi.Vars(r)

	greetingType := params["type"]

	switch greetingType {
	case "hello":
		fmt.Fprintln(w, "Hello World!!")
	case "goodbye":
		fmt.Fprintln(w, "Goodbye for now!!!")
	case "whatup":
		wagi.WriteDefaultErrorResponse(w, r)
	}
}

func main() {
	greetingHandler := wagi.InitHandlerWithRoute("/greeting/{type}").
		MethodHandler(http.MethodGet, DefRespHandler).
		WithDefaultResponse([]byte("So what's good?"))

	_ = cgi.Serve(greetingHandler)
}

package examples

import (
	"encoding/json"
	"fmt"
	"net/http"

	wagi "github.com/syke99/waggy/v2"
)

func ExampleWaggyHandler_WithDefaultResponse() {
	goodbyeHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello")
	}

	defaultResponse := make(map[string]interface{})

	defaultResponse["hello"] = "hello world"

	defaultResponseBytes, err := json.Marshal(defaultResponse)
	if err != nil {
		_ = fmt.Errorf("error marshaling body, err: %s", err)
	}

	handler := wagi.InitHandlerWithRoute("/example/test/route").
		MethodHandler(http.MethodDelete, goodbyeHandler).
		WithDefaultResponse(defaultResponseBytes)

	err = wagi.Serve(handler)
	if err != nil {
		_ = fmt.Errorf("error serving handler, err %s", err)
	}
}

func ExampleWaggyHandler_WithDefaultErrorResponse() {
	goodbyeHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello")
	}

	defaultError := wagi.WaggyError{
		Type:   "/example/test/route",
		Detail: "this request was bad",
		Status: http.StatusBadRequest,
	}

	handler := wagi.InitHandlerWithRoute("/example/test/route").
		MethodHandler(http.MethodDelete, goodbyeHandler).
		WithDefaultErrorResponse(defaultError, http.StatusBadRequest)

	err := wagi.Serve(handler)
	if err != nil {
		_ = fmt.Errorf("error serving handler, err %s", err)
	}
}

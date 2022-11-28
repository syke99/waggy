package waggy

// Init initializes the request and provides a *WaggyResponseWriter
// to use for writing responses, and a *WaggyRequest to use for
// retrieving info about the incoming HTTP request
func Init() (*WaggyResponseWriter, *WaggyRequest) {
	return Response(), Request()
}

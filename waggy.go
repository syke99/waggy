package waggy

// Init initializes the request and provides a *ResponseWriter
// to use for writing responses, and a *Request to use for
// retrieving info about the incoming HTTP request
func Init() (*ResponseWriter, *Request) {
	return Resp(), Req()
}

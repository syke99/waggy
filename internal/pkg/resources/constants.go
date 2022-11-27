package resources

type Env int

const (
	FullURL Env = iota
	PathInfo
	RawPathInfo
	PathTranslated
	Host
	Scheme
	Port
	Name
	QueryString
	RequestMethod
	RemoteAddr
	RemoteHost
	XMatchedRoute
	HTTPAccept
	HttpUserAgent
	ContentType
	ContentLength
	ScriptName
	ServerSoftware
	AuthType
	GatewayInterface
	RemoteUser
)

func (e Env) String() string {
	return []string{
		"X_FULL_URL",
		"PATH_INFO",
		"X_RAW_PATH_INFO",
		"PATH_TRANSLATED",
		"HTTP_HOST",
		"SERVER_PROTOCOL",
		"SERVER_PORT",
		"SERVER_NAME",
		"QUERY_STRING",
		"REQUEST_METHOD",
		"REMOTE_ADDR",
		"REMOTE_HOST",
		"X_MATCHED_ROUTE",
		"HTTP_ACCEPT",
		"HTTP_USER_AGENT",
		"CONTENT_TYPE",
		"CONTENT_LENGTH",
		"SCRIPT_NAME",
		"SERVER_SOFTWARE",
		"AUTH_TYPE",
		"GATEWAY_INTERFACE",
		"REMOTE_USER",
	}[e]
}

func GetStatusName(status int) string {
	name := ""

	switch status {
	case 100:
		name = "Continue"
	case 101:
		name = "Switching Protocols"
	case 103:
		name = "Early Hints"
	case 200:
		name = "OK"
	case 202:
		name = "Accepted"
	case 203:
		name = "Non Authoritative Information"
	case 204:
		name = "No Content"
	case 205:
		name = "Reset Content"
	case 206:
		name = "Partial Content"
	case 300:
		name = "Multiple Choices"
	case 301:
		name = "Moved Permanently"
	case 302:
		name = "Found"
	case 303:
		name = "See Other"
	case 304:
		name = "Not Modified"
	case 307:
		name = "Temporary Redirect"
	case 308:
		name = "Permanent Redirect"
	case 400:
		name = "Bad Request"
	case 401:
		name = "Unauthorized"
	case 402:
		name = "Payment Required"
	case 403:
		name = "Forbidden"
	case 404:
		name = "Not Found"
	case 405:
		name = "Method Not Allowed"
	case 406:
		name = "Not Acceptable"
	case 407:
		name = "Proxy Authentication Required"
	case 408:
		name = "Request Timed Out"
	case 409:
		name = "Conflict"
	case 410:
		name = "Gone"
	case 411:
		name = "Length Required"
	case 412:
		name = "Precondition Failed"
	case 413:
		name = "Payload Too Large"
	case 414:
		name = "URI Too Long"
	case 415:
		name = "unsupported Media Type"
	case 416:
		name = "Range Not Satisfied"
	case 417:
		name = "Expectation Failed"
	case 418:
		name = "I'm A Tea Pot"
	case 422:
		name = "Unprocessable Entity"
	case 425:
		name = "Too Early"
	case 426:
		name = "Upgrade Required"
	case 427:
		name = "Precondition Required"
	case 429:
		name = "Too Many Requests"
	case 431:
		name = "Request Header Fields Too Large"
	case 451:
		name = "Unavailable For Legal Reasons"
	case 500:
		name = "Internal Server Error"
	case 501:
		name = "Not Implemented"
	case 502:
		name = "Bad Gateway"
	case 503:
		name = "Service Unavailable"
	case 504:
		name = "Gateway Timeout"
	case 505:
		name = "HTTP Version Not Supported"
	case 506:
		name = "Variant Also Negotiates"
	case 507:
		name = "Insufficient Storage"
	case 508:
		name = "Loop Detected"
	case 510:
		name = "Not Extended"
	case 511:
		name = "Network Authentication Required"
	}

	return name
}

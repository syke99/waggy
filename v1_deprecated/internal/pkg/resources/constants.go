// v1 of Waggy has been deprecated. Please use v2 by running the following
// command: go get github.com/syke99/waggy/v2

package resources

//
//type Env int
//
//const (
//	FullURL Env = iota
//	PathInfo
//	RawPathInfo
//	PathTranslated
//	Host
//	Scheme
//	Port
//	Name
//	QueryString
//	RequestMethod
//	RemoteAddr
//	RemoteHost
//	XMatchedRoute
//	HTTPAccept
//	HttpUserAgent
//	ContentType
//	ContentLength
//	ScriptName
//	ServerSoftware
//	AuthType
//	GatewayInterface
//	RemoteUser
//)
//
//func (e Env) String() string {
//	return []string{
//		"X_FULL_URL",
//		"PATH_INFO",
//		"X_RAW_PATH_INFO",
//		"PATH_TRANSLATED",
//		"HTTP_HOST",
//		"SERVER_PROTOCOL",
//		"SERVER_PORT",
//		"SERVER_NAME",
//		"QUERY_STRING",
//		"REQUEST_METHOD",
//		"REMOTE_ADDR",
//		"REMOTE_HOST",
//		"X_MATCHED_ROUTE",
//		"HTTP_ACCEPT",
//		"HTTP_USER_AGENT",
//		"CONTENT_TYPE",
//		"CONTENT_LENGTH",
//		"SCRIPT_NAME",
//		"SERVER_SOFTWARE",
//		"AUTH_TYPE",
//		"GATEWAY_INTERFACE",
//		"REMOTE_USER",
//	}[e]
//}
//
//type StatusCode int
//
//var StatusCodes = map[StatusCode]string{
//	100: "Continue",
//	101: "Switching Protocols",
//	103: "Early Hints",
//	200: "OK",
//	202: "Accepted",
//	203: "Non Authoritative Information",
//	204: "No Content",
//	205: "Reset Content",
//	206: "Partial Content",
//	300: "Multiple Choices",
//	301: "Moved Permanently",
//	302: "Found",
//	303: "See Other",
//	304: "Not Modified",
//	307: "Temporary Redirect",
//	308: "Permanent Redirect",
//	400: "Bad Request",
//	401: "Unauthorized",
//	402: "Payment Required",
//	403: "Forbidden",
//	404: "Not Found",
//	405: "Method Not Allowed",
//	406: "Not Acceptable",
//	407: "Proxy Authentication Required",
//	408: "Request Timed Out",
//	409: "Conflict",
//	410: "Gone",
//	411: "Length Required",
//	412: "Precondition Failed",
//	413: "Payload Too Large",
//	414: "URI Too Long",
//	415: "unsupported Media Type",
//	416: "Range Not Satisfied",
//	417: "Expectation Failed",
//	418: "I'm A Tea Pot",
//	422: "Unprocessable Entity",
//	425: "Too Early",
//	426: "Upgrade Required",
//	427: "Precondition Required",
//	429: "Too Many Requests",
//	431: "Request Header Fields Too Large",
//	451: "Unavailable For Legal Reasons",
//	500: "Internal Server Error",
//	501: "Not Implemented",
//	502: "Bad Gateway",
//	503: "Service Unavailable",
//	504: "Gateway Timeout",
//	505: "HTTP Version Not Supported",
//	506: "Variant Also Negotiates",
//	507: "Insufficient Storage",
//	508: "Loop Detected",
//	510: "Not Extended",
//	511: "Network Authentication Required",
//}
//
//func (s StatusCode) GetStatusCodeName() string {
//	return StatusCodes[s]
//}

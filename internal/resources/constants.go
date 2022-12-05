package resources

type ContextKey int

const (
	DefResp ContextKey = iota
	DefErr
	PathParams
	QueryParams
)

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

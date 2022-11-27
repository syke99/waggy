package main

// Methods
const (
	MethodGet     = "GET"
	MethodPut     = "Put"
	MethodPost    = "POST"
	MethodPatch   = "PATCH"
	MethodDelete  = "DELETE"
	MethodTrace   = "TRACE"
	MethodHead    = "HEAD"
	MethodConnect = "CONNECT"
	MethodOptions = "OPTIONS"
)

// Status Codes
const (
	StatusContinue                      = 100
	StatusSwitchingProtocols            = 101
	StatusEarlyHints                    = 103
	StatusOK                            = 200
	StatusAccepted                      = 202
	StatusNonAuthoritativeInformation   = 203
	StatusNoContent                     = 204
	StatusResetContent                  = 205
	StatusPartialContent                = 206
	StatusMultipleChoices               = 300
	StatusMovedPermanently              = 301
	StatusFound                         = 302
	StatusSeeOther                      = 303
	StatusNotModified                   = 304
	StatusTemporaryRedirect             = 307
	StatusPermanentRedirect             = 308
	StatusBadRequest                    = 400
	StatusUnauthorized                  = 401
	StatusPaymentRequired               = 402
	StatusForbidden                     = 403
	StatusNotFound                      = 404
	StatusMethodNotAllowed              = 405
	StatusNotAcceptable                 = 406
	StatusProxyAuthenticationRequired   = 407
	StatusRequestTimedOut               = 408
	StatusConflict                      = 409
	StatusGone                          = 410
	StatusLengthRequired                = 411
	StatusPreconditionFailed            = 412
	StatusPayloadTooLarge               = 413
	StatusURITooLong                    = 414
	StatusUnsupportedMediaType          = 415
	StatusRangeNotSatisfiable           = 416
	StatusExpectationFailed             = 417
	StatusImATeaPot                     = 418
	StatusUnprocessableEntity           = 422
	StatusTooEarly                      = 425
	StatusUpgradeRequired               = 426
	StatusPreconditionRequired          = 427
	StatusTooManyRequests               = 429
	StatusRequestHeaderFieldsTooLarge   = 431
	StatusUnavailableForLegalReasons    = 451
	StatusInternalServerError           = 500
	StatusNotImplemented                = 501
	StatusBadGateway                    = 502
	StatusServiceUnavailable            = 503
	StatusGatewayTimeout                = 504
	StatusHTTPVersionNotSupported       = 505
	StatusVariantAlsoNegotiates         = 506
	StatusInsufficientStorage           = 507
	StatusLoopDetected                  = 508
	StatusNotExtended                   = 510
	StatusNetworkAuthenticationRequired = 511
)

// Headers
const (
	FullURL          = "X_FULL_URL"
	PathInfo         = "PATH_INFO"
	RawPathInfo      = "X_RAW_PATH_INFO"
	PathTranslated   = "PATH_TRANSLATED"
	Host             = "HTTP_HOST"
	Scheme           = "SERVER_PROTOCOL"
	Port             = "SERVER_PORT"
	Name             = "SERVER_NAME"
	QueryString      = "QUERY_STRING"
	RequestMethod    = "REQUEST_METHOD"
	RemoteAddr       = "REMOTE_ADDR"
	RemoteHost       = "REMOTE_HOST"
	XMatchedRoute    = "X_MATCHED_ROUTE"
	HTTPAccept       = "HTTP_ACCEPT"
	HttpUserAgent    = "HTTP_USER_AGENT"
	ContentType      = "CONTENT_TYPE"
	ContentLength    = "CONTENT_LENGTH"
	ScriptName       = "SCRIPT_NAME"
	ServerSoftware   = "SERVER_SOFTWARE"
	AuthType         = "AUTH_TYPE"
	GatewayInterface = "GATEWAY_INTERFACE"
	RemoteUser       = "REMOTE_USER"
)

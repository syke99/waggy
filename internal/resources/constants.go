package resources

type ContextKey int

const (
	DefResp ContextKey = iota
	DefErr
	MatchedRoute
	RootRoute
	PathParams
	QueryParams
)

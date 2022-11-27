package waggy

type WaggyRouter struct {
	moduleRoutes [][]Route
}

func NewWaggyRouter() *WaggyRouter {
	return &WaggyRouter{
		moduleRoutes: make([][]Route, 0),
	}
}

type Route struct {
	Path    string
	Method  string
	Handler func()
}

// this will map routes to handler functions in the
// modules.toml file
func (w *WaggyRouter) LoadRoutes(routes []Route) {
	w.moduleRoutes = append(w.moduleRoutes, routes)
}

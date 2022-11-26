package main

type Route struct {
	Path    string
	Method  string
	Handler func()
}

func LoadRoutes(...Route) {

}

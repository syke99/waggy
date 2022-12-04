package resources

import (
	"errors"
	"reflect"
	"runtime"
)

func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

const (
	Hello                     = "hello"
	Goodbye                   = "goodbye"
	WhereAmI                  = "where am I?"
	HelloWorld                = "hello world"
	TestFilePath              = "/test/file/path"
	TestRoute                 = "/test/route"
	TestRoutePathParams       = "/test/route/{param}"
	TestRoutePathParamHello   = "/test/route/hello"
	TestRoutePathParamGoodbye = "/test/route/goodbye"
)

var (
	TestError = errors.New("this is a test Waggy Error")
)

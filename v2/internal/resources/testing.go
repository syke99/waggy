package resources

import (
	"errors"
	"os"
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
	TestRoute                 = "/test/route"
	TestRoutePathParams       = "/test/route/{param}"
	TestRoutePathParamHello   = "/test/route/hello"
	TestRoutePathParamGoodbye = "/test/route/goodbye"
	HelloWorld                = "hello world"
)

var (
	TestError   = errors.New("this is a test Waggy Error")
	TestLogFile = &os.File{}
	TestKey     = "testKey"
	TestMessage = "testMessage"
	TestValue   = "testValue"
)

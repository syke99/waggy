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
	HelloWorld                = "hello world"
	TestFilePath              = "./internal/resources/testing.go"
	TestRoute                 = "/test/route"
	TestRoutePathParams       = "/test/route/{param}"
	TestRoutePathParamHello   = "/test/route/hello"
	TestRoutePathParamGoodbye = "/test/route/goodbye"
)

var (
	TestError   = errors.New("this is a test Waggy Error")
	TestLogFile = &os.File{}
	TestKey     = "testKey"
	TestMessage = "testMessage"
	TestValue   = "testValue"
)

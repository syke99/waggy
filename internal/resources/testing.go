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
	TestContentType           = "application/json"
	TestRoute                 = "/test/route"
	TestRoutePathParams       = "/test/route/{param}"
	TestRoutePathParamHello   = "/test/route/hello"
	TestRoutePathParamGoodbye = "/test/route/goodbye"
	TestErrorResponse         = "{ \"title\": \"Resource Not Found\", \"detail\": \"no path to file provided\", \"status\": \"404\" }"
)

var (
	TestError     = errors.New("this is a test Waggy Error")
	TestLogFile   = &os.File{}
	TestKey       = "testKey"
	TestMapKey1   = "testMapKey1"
	TestMapKey2   = "testMapKey2"
	TestMapKey3   = "testMapKey3"
	TestMapValue1 = []string{Hello}
	TestMapValue2 = []string{Hello, Goodbye}
	TestMapValue3 = make([]string, 0)
	TestMessage   = "testMessage"
	TestValue     = "testValue"
	TestQueryMap  = func() map[string][]string {
		m := make(map[string][]string)
		m[TestMapKey1] = TestMapValue1
		m[TestMapKey2] = TestMapValue2
		m[TestMapKey3] = TestMapValue3

		return m
	}
	TestEnvs = []string{}
)

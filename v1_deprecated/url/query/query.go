// v1 of Waggy has been deprecated. Please use v2 by running the following
// command: go get github.com/syke99/waggy/v2

package query

//
//import (
//	"net/url"
//	"os"
//
//	"github.com/syke99/waggy/v1/internal/pkg/resources"
//)
//
//// Query returns the parsed query values
//type Query struct {
//	queryStr  string
//	queryArgs url.Values
//}
//
//// GetQuery parses query values and returns a new Query struct containing these values
//func GetQuery() *Query {
//	queryStr := os.Getenv(resources.QueryString.String())
//
//	queryArgs, _ := url.ParseQuery(queryStr)
//
//	q := Query{
//		queryStr:  queryStr,
//		queryArgs: queryArgs,
//	}
//
//	return &q
//}
//
//// Get works just like URL.Query().Get() in net/url
//func (q *Query) Get(key string) string {
//	return q.queryArgs.Get(key)
//}
//
//// Has works just like URL.Query().Has() in net/url
//func (q *Query) Has(key string) bool {
//	return q.queryArgs.Has(key)
//}
//
//// Set works just like URL.Query().Set() in net/url
//func (q *Query) Set(key string, value string) {
//	q.queryArgs.Set(key, value)
//}
//
//// Del works just like URL.Query().Del() in net/url
//func (q *Query) Del(key string) {
//	q.queryArgs.Del(key)
//}

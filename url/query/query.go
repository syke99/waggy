package query

import (
	"net/url"
	"os"

	"github.com/syke99/waggy/internal/pkg/resources"
)

type Query struct {
	queryStr  string
	queryArgs url.Values
}

func GetQuery() *Query {

	queryStr := os.Getenv(resources.QueryString.String())

	queryArgs, _ := url.ParseQuery(queryStr)

	q := Query{
		queryStr:  queryStr,
		queryArgs: queryArgs,
	}

	return &q
}

func (q *Query) Get(key string) string {
	return q.queryArgs.Get(key)
}

func (q *Query) Has(key string) bool {
	return q.queryArgs.Has(key)
}

func (q *Query) Set(key string, value string) {
	q.queryArgs.Set(key, value)
}

func (q *Query) Del(key string) {
	q.queryArgs.Del(key)
}

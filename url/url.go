package url

import (
	"github.com/syke99/waggy/internal/pkg/resources"
	"github.com/syke99/waggy/url/query"
	"os"
	"strconv"
)

type URL struct {
	fullUrl string
	host    string
	scheme  string
	port    string
}

func GetUrl() *URL {

	u := URL{
		fullUrl: os.Getenv(resources.FullURL),
		host:    os.Getenv(resources.Host),
		scheme:  os.Getenv(resources.Scheme),
		port:    os.Getenv(resources.Port),
	}

	return &u
}

func (u *URL) Query() *query.Query {
	return query.GetQuery()
}

func (u *URL) RawQuery() string {
	return os.Getenv(resources.QueryString)
}

func (u *URL) Host() string {
	return u.host
}

func (u *URL) Scheme() string {
	return u.scheme
}

func (u *URL) Port() int {
	p, _ := strconv.Atoi(u.port)

	return p
}

func (u *URL) Path() string {
	return u.fullUrl
}

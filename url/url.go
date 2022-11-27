package url

import (
	"os"
	"strconv"

	"github.com/syke99/waggy/internal/pkg/resources"
	"github.com/syke99/waggy/url/query"
)

type URL struct {
	fullUrl     string
	pathInfo    string
	rawPathInfo string
	host        string
	scheme      string
	port        string
	query       *query.Query
}

func GetUrl() *URL {

	u := URL{
		fullUrl:     os.Getenv(resources.FullURL.String()),
		pathInfo:    os.Getenv(resources.PathInfo.String()),
		rawPathInfo: os.Getenv(resources.RawPathInfo.String()),
		host:        os.Getenv(resources.Host.String()),
		scheme:      os.Getenv(resources.Scheme.String()),
		port:        os.Getenv(resources.Port.String()),
		query:       query.GetQuery(),
	}

	return &u
}

func (u *URL) RawQuery() string {
	return os.Getenv(resources.QueryString.String())
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
	return u.pathInfo
}

func (u *URL) RawPath() string {
	return u.rawPathInfo
}

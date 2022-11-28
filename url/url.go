package url

import (
	"os"
	"strconv"

	"github.com/syke99/waggy/internal/pkg/resources"
	"github.com/syke99/waggy/url/query"
)

// URL used for accessing URL specific WaggyRequest values
type URL struct {
	fullUrl     string
	pathInfo    string
	rawPathInfo string
	host        string
	scheme      string
	port        string
	Query       *query.Query
}

// GetUrl loads all URL-pertinent values into a new URL struct and returns it
func GetUrl() *URL {
	u := URL{
		fullUrl:     os.Getenv(resources.FullURL.String()),
		pathInfo:    os.Getenv(resources.PathInfo.String()),
		rawPathInfo: os.Getenv(resources.RawPathInfo.String()),
		host:        os.Getenv(resources.Host.String()),
		scheme:      os.Getenv(resources.Scheme.String()),
		port:        os.Getenv(resources.Port.String()),
		Query:       query.GetQuery(),
	}

	return &u
}

// RawQuery returns the query string without being url-decoded
func (u *URL) RawQuery() string {
	return os.Getenv(resources.QueryString.String())
}

// Host returns the value of the client-supplied HOST header
func (u *URL) Host() string {
	return u.host
}

// Scheme returns the protocol that the server is using. Usually it is HTTP/1.1
func (u *URL) Scheme() string {
	return u.scheme
}

// Port returns the port upon which the server received its request
func (u *URL) Port() int {
	p, _ := strconv.Atoi(u.port)

	return p
}

// Path returns a WAGI-specific representation of the Path;
// a full explanation can be found here ( https://github.com/deislabs/wagi/blob/main/docs/environment_variables.md )
func (u *URL) Path() string {
	return u.pathInfo
}

// RawPath returns a WAGI-specific representation of the raw, non url-decoded Path;
// a full explanation can be found here ( https://github.com/deislabs/wagi/blob/main/docs/environment_variables.md )
func (u *URL) RawPath() string {
	return u.rawPathInfo
}

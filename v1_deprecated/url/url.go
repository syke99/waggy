// v1 of Waggy has been deprecated. Please use v2 by running the following
// command: go get github.com/syke99/waggy/v2

package url

//
//import (
//	"github.com/syke99/waggy/v1/internal/pkg/resources"
//	"github.com/syke99/waggy/v1/url/query"
//	"os"
//	"strconv"
//	"strings"
//)
//
//// URL used for accessing URL specific Request values
//type URL struct {
//	fullUrl     string
//	pathInfo    string
//	rawPathInfo string
//	Params      map[string]string
//	host        string
//	scheme      string
//	port        string
//	query       *query.Query
//}
//
//// GetUrl loads all URL-pertinent values into a new URL struct and returns it
//func GetUrl(pathParams map[int]string) *URL {
//
//	u := URL{
//		fullUrl:     os.Getenv(resources.FullURL.String()),
//		pathInfo:    os.Getenv(resources.PathInfo.String()),
//		rawPathInfo: os.Getenv(resources.RawPathInfo.String()),
//		Params:      make(map[string]string),
//		host:        os.Getenv(resources.Host.String()),
//		scheme:      os.Getenv(resources.Scheme.String()),
//		port:        os.Getenv(resources.Port.String()),
//		query:       query.GetQuery(),
//	}
//
//	if len(pathParams) != 0 {
//		splitRawPath := strings.Split(u.pathInfo, "/")
//
//		params := make(map[string]string)
//
//		for k, v := range pathParams {
//			if v[:1] == "{" &&
//				v[len(v)-1:] == "}" {
//				params[splitRawPath[k]] = v[1 : len(v)-1]
//			}
//		}
//
//		u.Params = params
//	}
//
//	return &u
//}
//
//func (u *URL) Query() *query.Query {
//	return u.query
//}
//
//// RawQuery returns the query string without being url-decoded
//func (u *URL) RawQuery() string {
//	return os.Getenv(resources.QueryString.String())
//}
//
//// Host returns the value of the client-supplied HOST header
//func (u *URL) Host() string {
//	return u.host
//}
//
//// Scheme returns the protocol that the server is using. Usually it is HTTP/1.1
//func (u *URL) Scheme() string {
//	return u.scheme
//}
//
//// Port returns the port upon which the server received its request
//func (u *URL) Port() int {
//	p, _ := strconv.Atoi(u.port)
//
//	return p
//}
//
//// Path returns a WAGI-specific representation of the Path;
//// a full explanation can be found here ( https://github.com/deislabs/wagi/blob/main/docs/environment_variables.md )
//func (u *URL) Path() string {
//	return u.pathInfo
//}
//
//// RawPath returns a WAGI-specific representation of the raw, non url-decoded Path;
//// a full explanation can be found here ( https://github.com/deislabs/wagi/blob/main/docs/environment_variables.md )
//func (u *URL) RawPath() string {
//	return u.rawPathInfo
//}

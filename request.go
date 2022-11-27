package waggy

import (
	"github.com/syke99/waggy/header"
	"io"
	"io/ioutil"
	"os"

	"github.com/syke99/waggy/internal/pkg/resources"
	"github.com/syke99/waggy/url"
)

// WaggyRequest used for accessing information about the specific HTTP Request made
type WaggyRequest struct {
	Body       io.Reader
	URL        *url.URL
	Header     *header.Header
	method     string
	remoteAddr string
}

// Request loads the incoming HTTP Request into a new WaggyRequest struct
func Request() *WaggyRequest {

	wr := WaggyRequest{
		Body:       os.Stdin,
		URL:        url.GetUrl(),
		Header:     header.GetHeaders(),
		method:     os.Getenv(resources.RequestMethod.String()),
		remoteAddr: os.Getenv(resources.RemoteAddr.String()),
	}

	return &wr
}

// GetBody returns a slice of bytes read from the WaggyRequest's Body
func (r *WaggyRequest) GetBody() ([]byte, error) {
	return ioutil.ReadAll(r.Body)
}

// Method returns the HTTP Method used in the specific WaggyRequest
func (r *WaggyRequest) Method() string {
	return r.method
}

// RemoteAddr returns the client's IP address
func (r *WaggyRequest) RemoteAddr() string {
	return r.remoteAddr
}

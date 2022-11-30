package waggy

import (
	"bufio"
	"bytes"
	"github.com/syke99/waggy/header"
	"github.com/syke99/waggy/mime"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/syke99/waggy/internal/pkg/resources"
	"github.com/syke99/waggy/url"
)

// Request used for accessing information about the specific HTTP Request made
type Request struct {
	body          io.Reader
	MultipartForm *mime.MultipartForm
	URL           *url.URL
	Header        *header.Header
	method        string
	remoteAddr    string
}

// Reqloads the incoming HTTP Request into a new Request struct
func Req(opt RouteOption) *Request {
	wr := Request{
		body:          os.Stdin,
		MultipartForm: mime.GetMultipartForm(),
		URL:           url.GetUrl(opt.pathParams),
		Header:        header.GetHeaders(),
		method:        os.Getenv(resources.RequestMethod.String()),
		remoteAddr:    os.Getenv(resources.RemoteAddr.String()),
	}

	return &wr
}

// GetBody returns a slice of bytes read from the Request's Body
func (r *Request) GetBody() ([]byte, error) {
	return ioutil.ReadAll(r.body)
}

// Method returns the HTTP Method used in the specific Request
func (r *Request) Method() string {
	return r.method
}

// RemoteAddr returns the client's IP address
func (r *Request) RemoteAddr() string {
	return r.remoteAddr
}

// ParseMultipartForm parses the Request's Body as a multipart form and stores each form part
// in a map that is stored in r.MultipartForm. Each form part is stored at a key corresponding to the
// value supplied in the name portion of the form part's Content-Disposition header
func (r *Request) ParseMultipartForm() error {
	contentTypeHeaders := r.Header.Values("Content-Type")

	boundary := ""

	for _, value := range contentTypeHeaders {
		// attempt to split
		splitValue := strings.Split(value, "=")

		if splitValue[0] == "boundary" {
			boundary = splitValue[1]
			break
		}
	}

	body, err := r.GetBody()

	if err == nil {
		formParts := bytes.Split(body, []byte(boundary))

		for _, formPart := range formParts {
			buf := bytes.NewBuffer(formPart)

			scanner := bufio.NewScanner(buf)

			scanner.Scan()

			name := strings.Split(scanner.Text(), " ")[1]

			r.MultipartForm.Set(name, formPart)
			continue
		}
	}

	return err
}

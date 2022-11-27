package main

import (
	"github.com/syke99/waggy/header"
	"io"
	"io/ioutil"
	"os"

	"github.com/syke99/waggy/internal/pkg/resources"
	"github.com/syke99/waggy/url"
)

type WaggyRequest struct {
	Body       io.Reader
	URL        *url.URL
	Header     *header.Header
	method     string
	remoteAddr string
}

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

func (r *WaggyRequest) GetBody() ([]byte, error) {
	return ioutil.ReadAll(r.Body)
}

func (r *WaggyRequest) Method() string {
	return r.method
}

func (r *WaggyRequest) RemoteAddr() string {
	return r.remoteAddr
}

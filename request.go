package main

import (
	"github.com/syke99/waggy/internal/pkg/resources"
	"github.com/syke99/waggy/url"
	"io"
	"io/ioutil"
	"os"
)

type WaggyRequest struct {
	body       io.Reader
	method     string
	remoteAddr string
}

func Request() *WaggyRequest {
	wr := WaggyRequest{
		body:       os.Stdin,
		method:     os.Getenv(resources.RequestMethod),
		remoteAddr: os.Getenv(resources.RemoteAddr),
	}

	return &wr
}

func (r *WaggyRequest) GetBody() ([]byte, error) {
	return ioutil.ReadAll(r.body)
}

func (r *WaggyRequest) URL() *url.URL {
	return url.GetUrl()
}

func (r *WaggyRequest) Method() string {
	return r.method
}

func (r *WaggyRequest) RemoteAddr() string {
	return r.remoteAddr
}

package fetch

import (
	"io"
	"net/http"
)

type RequestInit struct {
	Method string
	Header http.Header
	Body   io.Reader
}

func Run(client *http.Client, addr string, opt *RequestInit) (res *http.Response, err error) {
	req, err := http.NewRequest(opt.Method, addr, opt.Body)
	if err != nil {
		return
	}
	copyHeader(req.Header, opt.Header)
	res, err = client.Do(req)
	return
}

func copyHeader(dst http.Header, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}

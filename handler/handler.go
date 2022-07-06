package handler

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/shynome/httprelay-go/fetch"
)

type RequestInit func(req *http.Request) *http.Request

type Handler struct {
	Handler http.Handler
	Secret  string
}

func MakeRequestInit(secret string, header http.Header, body io.Reader) *fetch.RequestInit {
	if secret != "" {
		header.Set("HttpRelay-WSecret", secret)
	}
	return &fetch.RequestInit{
		Method: "SERVE",
		Header: header,
		Body:   body,
	}
}

func (h *Handler) Execute(ctx Context) *fetch.RequestInit {
	var w = httptest.NewRecorder()
	h.Handler.ServeHTTP(w, ctx.Request)
	res := w.Result()
	res = NewResponse(res)

	res.Header.Set("HttpRelay-Proxy-JobId", ctx.JobId())
	init := MakeRequestInit(ctx.Secret, res.Header, res.Body)
	return init
}

func NewRequest(res *http.Response) (req *http.Request, err error) {
	var (
		method string
		link   string
	)
	header := res.Header
	if method, err = headerValue(header, "HttpRelay-Proxy-Method"); err != nil {
		return
	}
	if link, err = headerValue(header, "HttpRelay-Proxy-Url"); err != nil {
		return
	}
	req, err = http.NewRequest(method, link, nil)
	if err != nil {
		return
	}
	if req.URL.Path, err = headerValue(header, "HttpRelay-Proxy-Path"); err != nil {
		return
	}
	req.URL.RawQuery = header.Get("HttpRelay-Proxy-Query")

	req.RequestURI = req.URL.Path
	if req.URL.RawQuery != "" {
		req.RequestURI += "?" + req.URL.Query().Encode()
	}
	for k, vv := range header {
		for _, v := range vv {
			req.Header.Add(k, v)
		}
	}
	req.Body = res.Body
	return
}

func headerValue(header http.Header, k string) (v string, err error) {
	v = header.Get(k)
	if v == "" {
		err = fmt.Errorf("unable to find \"%s\" header field", k)
	}
	return
}

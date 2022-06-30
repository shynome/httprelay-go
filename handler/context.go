package handler

import "net/http"

type Context struct {
	Request *http.Request
	Secret  string
}

func (ctx *Context) JobId() string {
	return ctx.Request.Header.Get("HttpRelay-Proxy-JobId")
}

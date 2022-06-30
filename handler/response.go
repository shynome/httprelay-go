package handler

import (
	"fmt"
	"net/http"
	"strings"
)

func NewResponse(resp *http.Response) *http.Response {
	header := resp.Header

	headerWhiteList := func() string {
		headerWhiteList := []string{}
		for k := range header {
			headerWhiteList = append(headerWhiteList, k)
		}
		return strings.Join(headerWhiteList, ", ")
	}()

	header.Set("httprelay-proxy-headers", headerWhiteList)
	header.Set("httprelay-proxy-status", fmt.Sprintf("%d", resp.StatusCode))

	return resp
}

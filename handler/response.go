package handler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/shynome/httprelay-go/fetch"
)

func NewResponse(resp *http.Response) *http.Response {
	header := resp.Header

	headerWhiteList := func() string {
		headerWhiteList := []string{}
		for k := range fetch.WriteExcludeHeaderDump {
			val := header.Get(k)
			if len(val) == 0 {
				continue
			}
			header.Set("X-Pass-Dump-"+k, header.Get(k))
		}
		for k := range header {
			if fetch.WriteExcludeHeaderDump[k] {
				continue
			}
			headerWhiteList = append(headerWhiteList, k)
		}
		return strings.Join(headerWhiteList, ", ")
	}()

	header.Set("httprelay-proxy-headers", headerWhiteList)
	header.Set("httprelay-proxy-status", fmt.Sprintf("%d", resp.StatusCode))

	return resp
}

package fetch

// https://github.com/golang/go/blob/26049f6f9171d1190f3bbe05ec304845cfe6399f/src/net/http/httputil/dump.go#L164
var WriteExcludeHeaderDump = map[string]bool{
	"Host":              true, // not in Header map anyway
	"Content-Length":    true,
	"Transfer-Encoding": true,
	"Trailer":           true,
}

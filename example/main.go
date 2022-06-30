package main

import (
	"fmt"
	"net/http"

	"github.com/shynome/httprelay-go"
)

func main() {

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello world")
	})

	proxy := httprelay.NewProxy("http://127.0.0.1:8080")
	proxy.Auth = httprelay.ProxyAuth{ID: "aaaaa", Secret: "88888888"}
	proxy.Parallel = 1

	fmt.Printf("%shello \n", proxy.GetServerUrl())
	proxy.Serve(nil)
}

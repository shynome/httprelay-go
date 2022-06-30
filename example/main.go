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
	server := "http://127.0.0.1:8080"
	auth := httprelay.ProxyAuth{ID: "aaaaa", Secret: "88888888"}
	relay := httprelay.NewProxy(server)
	// relay.Parallel = 1
	relay.Auth = auth
	fmt.Printf("%shello \n", relay.GetServerUrl())
	relay.Serve(nil)
}

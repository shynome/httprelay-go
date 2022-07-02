# Description

[httprelay](https://gitlab.com/jonas.jasas/httprelay) golang client,
like [httprelay-js](https://gitlab.com/jonas.jasas/httprelay-js), but only impl proxy

# Usage

```go
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

	proxy := httprelay.NewProxy("https://demo.httprelay.io")

	fmt.Printf("%shello \n", proxy.GetServerUrl())
	proxy.Serve(nil)
}

```

# Benchmark

```sh
docker run --net host --rm -i grafana/k6 run - <k6.js
```

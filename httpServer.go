// http server
package main

import (
	"fmt"
	"net/http"
)

func handler(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(rw, "hello world!")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8081", nil)
}

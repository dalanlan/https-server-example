// https server

package main

import (
	"fmt"
	"net/http"
)

func handler(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(rw, "hello world, https!")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServeTLS(":6669", "serverBi.crt", "serverBi.key", nil)
}

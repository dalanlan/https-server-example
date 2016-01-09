// https server which verifies client cert

package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
)

type myHandler struct {
}

func (h *myHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(rw, "hello https, 9999!")
}

func main() {
	pool := x509.NewCertPool()
	caCrt, err := ioutil.ReadFile("ca.crt")
	if err != nil {
		fmt.Println(err)
		return
	}

	pool.AppendCertsFromPEM(caCrt)

	s := &http.Server{
		Addr:    ":9999",
		Handler: &myHandler{},
		TLSConfig: &tls.Config{
			ClientCAs:  pool,
			ClientAuth: tls.RequireAndVerifyClientCert,
		},
	}

	err = s.ListenAndServeTLS("serverBi.crt", "serverBi.key")
	if err != nil {
		fmt.Println(err)
	}
}

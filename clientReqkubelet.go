// client request for kubelet server

package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	pool := x509.NewCertPool()
	caCrt, err := ioutil.ReadFile("/home/vcap/kubelet/ca.crt")
	if err != nil {
		fmt.Println(err)
		return
	}
	pool.AppendCertsFromPEM(caCrt)

	cliCrt, err := tls.LoadX509KeyPair("/home/vcap/kubelet/client.crt", "/home/vcap/kubelet/client.key")
	if err != nil {
		fmt.Println(err)
		return
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			RootCAs:      pool,
			Certificates: []tls.Certificate{cliCrt},
		},
	}
	client := &http.Client{Transport: tr}

	resp, err := client.Get("https://localhost:10250")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

}

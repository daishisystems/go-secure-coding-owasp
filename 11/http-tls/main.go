package main

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
)

type Certificates struct {
	CertFile string
	KeyFile  string
}

func main() {
	httpsServer := &http.Server{
		Addr: ":8080",
	}
	var certs []Certificates
	certs = append(certs, Certificates{
		CertFile: "../etc/yourSite.pem", //Your site certificate key
		KeyFile:  "../etc/yourSite.key", //Your site private key
	})
	config := &tls.Config{}

	config.Certificates = make([]tls.Certificate, len(certs))
	for i, v := range certs {
		config.Certificates[i], _ = tls.LoadX509KeyPair(v.CertFile, v.KeyFile)
	}
	conn, _ := net.Listen("tcp", ":8080")
	tlsListener := tls.NewListener(conn, config)
	httpsServer.Serve(tlsListener)
	fmt.Println("Listening on port 8080...")
}

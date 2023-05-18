package httputils

import (
	"crypto/tls"
	"net/http"
)

func CreateCustomHTTPClient(insecureSkipVerify bool, serverName string) *http.Client {
	customTransport := http.DefaultTransport.(*http.Transport).Clone()
	customTransport.TLSClientConfig = &tls.Config{
		InsecureSkipVerify: insecureSkipVerify,
		ServerName:         serverName,
	}
	client := &http.Client{Transport: customTransport}
	return client
}

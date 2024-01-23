package http

import (
	"crypto/tls"
	"net/http"
	"time"
)

func NewHttpClient() *http.Client {
	client := &http.Client{
		Timeout: 2 * time.Minute,
		Transport: &http.Transport{
			MaxIdleConns:        1024,
			MaxConnsPerHost:     100,
			MaxIdleConnsPerHost: 100,
			TLSClientConfig: &tls.Config{
				MinVersion:         1.0,
				InsecureSkipVerify: true,
				Renegotiation:      tls.RenegotiateOnceAsClient,
			},
		},
	}
	return client
}

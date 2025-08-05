package config

import (
	"net"
	"net/http"
	"time"
)

var HttpClient *http.Client

func InitHTTPClient() {
	transport := &http.Transport{
		MaxIdleConns:        500,
		MaxIdleConnsPerHost: 100,

		IdleConnTimeout:     90 * time.Second,
		TLSHandshakeTimeout: 10 * time.Second,
		DialContext: (&net.Dialer{
			Timeout:   5 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
	}

	HttpClient = &http.Client{
		Transport: transport,
		Timeout:   60 * time.Second,
	}
}

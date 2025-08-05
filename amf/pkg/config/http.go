package config

import (
	"crypto/tls"
	"net"
	"net/http"
	"time"

	"golang.org/x/net/http2"
)

const NUM_CLIENT = 32

var ListHttpClient [NUM_CLIENT]*http.Client

func InitHTTPClient() {
	for i := 0; i < NUM_CLIENT; i++ {
		transport := &http2.Transport{
			AllowHTTP:                  true,
			StrictMaxConcurrentStreams: false,
			DisableCompression:         true,
			IdleConnTimeout:            10 * time.Second,
			DialTLS: func(network, addr string, cfg *tls.Config) (net.Conn, error) {
				return net.DialTimeout(network, addr, 3*time.Second)
			},
		}

		httpClient := &http.Client{
			Transport: transport,
			Timeout:   1 * time.Second,
		}

		ListHttpClient[i] = httpClient
	}
}

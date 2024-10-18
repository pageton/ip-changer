package proxy

import (
	"net/http"

	"golang.org/x/net/proxy"
)

func NewHTTPClient() *http.Client {
	dialer, _ := proxy.SOCKS5("tcp", "127.0.0.1:9050", nil, proxy.Direct)
	transport := &http.Transport{Dial: dialer.Dial}
	return &http.Client{Transport: transport}
}

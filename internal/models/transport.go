package models

import (
	"examtopics-downloader/internal/constants"
	"net"
	"net/http"
	"time"
)

// OptimizedTransport returns a high-performance HTTP transport
func OptimizedTransport() *http.Transport {
	dialer := &net.Dialer{
		Timeout:       10 * time.Second,
		KeepAlive:     30 * time.Second,
		FallbackDelay: 250 * time.Millisecond,
	}
	return &http.Transport{
		Proxy:                 http.ProxyFromEnvironment,
		DialContext:           dialer.DialContext,
		MaxIdleConns:          constants.MaxIdleConns,
		MaxIdleConnsPerHost:   constants.MaxIdleConnsPerHost,
		MaxConnsPerHost:       constants.MaxConnsPerHost,
		IdleConnTimeout:       constants.IdleConnTimeout,
		DisableCompression:    false,
		DisableKeepAlives:     false,
		TLSHandshakeTimeout:   constants.TLSHandshakeTimeout,
		ResponseHeaderTimeout: constants.ResponseHeaderTimeout,
		ExpectContinueTimeout: constants.ExpectContinueTimeout,
	}
}

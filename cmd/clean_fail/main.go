package main

import (
	"http_keepalive/lib"
	"net"
	"net/http"
	"time"
)

func main() {
	client := http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
				DualStack: true,
			}).DialContext,
			MaxIdleConns:          100,
			MaxIdleConnsPerHost:   100, // TRICK!
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
	}

	go func() {
		time.Sleep(time.Second * 2)
		for i := 0; i < 1000; i++ {
			go lib.RequestClean(&client)
		}
	}()

	lib.Serve()
}

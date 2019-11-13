package main

import (
	"http_keepalive/lib"
	"net/http"
	"time"
)

func main() {
	myTransport := *http.DefaultTransport.(*http.Transport)
	myTransport.MaxIdleConnsPerHost = 100 // TRICK!
	myTransport.MaxConnsPerHost = 100 // TRICKIER!
	client := http.Client{
		Transport: &myTransport,
	}

	go func() {
		time.Sleep(time.Second * 2)
		for i := 0; i < 1000; i++ {
			go lib.RequestClean(&client)
		}
	}()

	lib.Serve()
}

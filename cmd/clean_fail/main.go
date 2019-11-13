package main

import (
	"http_keepalive/lib"
	"net/http"
	"time"
)

func main() {
	go func() {
		time.Sleep(time.Second * 2)
		for i := 0; i < 1000; i++ {
			go lib.RequestClean(http.DefaultClient)
			time.Sleep(time.Millisecond * 1)
		}
	}()
	
	lib.Serve()
}

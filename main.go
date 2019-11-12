package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

func serve() {
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		if _, err := writer.Write([]byte("world")); err != nil {
			panic(err)
		}
	})
	if err := http.ListenAndServe("0.0.0.0:15000", nil); err != nil {
		panic(err)
	}
}

func requestWithoutReadAndClose(client *http.Client) {
	_, err := client.Get("http://localhost:15000/hello")
	if err != nil {
		panic(err)
	}
}

func requestWithoutRead(client *http.Client) {
	resp, err := client.Get("http://localhost:15000/hello")
	if err != nil {
		panic(err)
	}

	if err := resp.Body.Close(); err != nil {
		panic(err)
	}
}

func requestWithoutClose(client *http.Client) {
	resp, err := client.Get("http://localhost:15000/hello")
	if err != nil {
		panic(err)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}

func requestClean(client *http.Client) {
	resp, err := client.Get("http://localhost:15000/hello")
	if err != nil {
		panic(err)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
	if err := resp.Body.Close(); err != nil {
		panic(err)
	}
}

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
			MaxIdleConnsPerHost:   100,
			MaxConnsPerHost:       100,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
	}

	go func() {
		time.Sleep(time.Second * 2)
		for i := 0; i < 1000; i++ {
			go requestClean(&client)
			// time.Sleep(time.Millisecond * 10)
		}
	}()

	serve()
}

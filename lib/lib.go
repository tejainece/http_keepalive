package lib

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Serve() {
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		if _, err := writer.Write([]byte("world")); err != nil {
			panic(err)
		}
	})
	if err := http.ListenAndServe("0.0.0.0:15000", nil); err != nil {
		panic(err)
	}
}

func RequestWithoutReadAndClose(client *http.Client) {
	_, err := client.Get("http://localhost:15000/hello")
	if err != nil {
		panic(err)
	}
}

func RequestWithoutRead(client *http.Client) {
	resp, err := client.Get("http://localhost:15000/hello")
	if err != nil {
		panic(err)
	}

	if err := resp.Body.Close(); err != nil {
		panic(err)
	}
}

func RequestWithoutClose(client *http.Client) {
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

func RequestClean(client *http.Client) {
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

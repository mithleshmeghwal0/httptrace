package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptrace"
	"time"
)

var clientTrace *httptrace.ClientTrace

func init() {
	clientTrace = &httptrace.ClientTrace{
		GetConn: func(hostPort string) {
			fmt.Println("starting to create conn ", hostPort)
		},
		DNSStart: func(info httptrace.DNSStartInfo) {
			fmt.Println("starting to look up dns", toString(info))
		},
		DNSDone: func(info httptrace.DNSDoneInfo) {
			fmt.Println("done looking up dns", toString(info))
		},
		ConnectStart: func(network string, addr string) {
			fmt.Println("starting tcp connection", toString(map[string]interface{}{"network": network, "addr": addr}))
		},
		ConnectDone: func(network string, addr string, err error) {
			fmt.Println("tcp connection created", toString(map[string]interface{}{"network": network, "addr": addr, "err": err}))
		},
		GotConn: func(info httptrace.GotConnInfo) {
			fmt.Println("connection established", toString(info))
		},
	}

}

var routes = []string{
	"https://www.bing.com/search?q=",
	"https://www.google.com/search?q=",
}

var keywords = []string{
	"ninja",
	"cat",
	"dog",
	"golang",
	"html",
	"nasa",
	"tissue",
	"coffee",
}

func main() {
	path := routes[0]
	var client http.Client
	for i := 0; i < 8; i++ {
		//changeSearchEngine()
		if i == 3 {
			path = routes[1]
		}
		if i == 6 {
			path = routes[0]
		}
		req, err := http.NewRequest(http.MethodGet, path+keywords[i], nil)
		if err != nil {
			panic(err)
		}

		clientTraceCtx := httptrace.WithClientTrace(req.Context(), clientTrace)
		req = req.WithContext(clientTraceCtx)
		_, err = client.Do(req)
		if err != nil {
			panic(err)
		}
		fmt.Println()
		time.Sleep(2 * time.Second)
	}
}

func toString(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}

	return string(b)
}

package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func main() {
	values := url.Values{"test": {"value"}}
	reader := strings.NewReader(values.Encode())

	client := &http.Client{}
	request, err := http.NewRequest("DELETE", "http://localhost:18888", reader)
	if err != nil {
		panic(err)
	}
	resp, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	dump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	log.Println(string(dump))
}

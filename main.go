package main

import (
	"fmt"
	"net/http"
)

type realImplementation struct {
	httpClient http.Client
}

func (r realImplementation) getHttpClient() {
	r.httpClient = http.Client{}
}

func main() {
	r := realImplementation{}
	r.getHttpClient()

	resp, _ := r.httpClient.Get("https://google.com")
	fmt.Println(resp)
}

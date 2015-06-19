package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {
	var response *http.Response
	var request *http.Request

	url := "http://localhost/foo/bar"
	request, err := http.NewRequest("GET", url, nil)
	if err == nil {
		request.Header.Add("Content-Type", "application/json")
		debug(httputil.DumpRequestOut(request, true))
		response, err = (&http.Client{}).Do(request)
	}

	if err == nil {
		defer response.Body.Close()
		debug(httputil.DumpResponse(response, true))
	}
}

func debug(data []byte, err error) {
	if err == nil {
		fmt.Printf("%s\n\n", data)
	} else {
		log.Fatalf("%s\n\n", err)
	}
}

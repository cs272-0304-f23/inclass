package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
)

type TestCase struct {
	expected string
}

func main() {
	// could use an anonymous struct and fill in test case data
	tests := []TestCase{
		{expected: "my response"},
	}

	for _, tc := range tests {
		// create a mock http server which responds with MyHandler
		// anonymous function literals can be used when a function name is expected
		svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("my response"))
		}))
		defer svr.Close()

		// act like an HTTP client to download the document from
		// the URL and compare it to what we expect to get
		if rsp, err := http.Get(svr.URL); err == nil {
			if bts, err := io.ReadAll(rsp.Body); err == nil {
				if bytes.Equal(bts, []byte(tc.expected)) {
					fmt.Println("pass!")
				}
			}
		}
	}
}

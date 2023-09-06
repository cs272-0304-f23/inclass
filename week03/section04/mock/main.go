package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
)

type TestCase struct {
	expected []byte
}

func main() {
	tests := []TestCase{
		{expected: []byte("good afternoon")},
	}

	for _, tc := range tests {
		// Launch mock http server
		svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("good afternoon"))
		}))
		defer svr.Close()

		// Use http client to get the mock server's URL and port
		if rsp, err := http.Get(svr.URL); err == nil {
			if bts, err := io.ReadAll(rsp.Body); err == nil {
				if bytes.Equal(bts, tc.expected) {
					fmt.Println("pass")
				} else {
					fmt.Println("fail")
				}
			}
		}
	}
}

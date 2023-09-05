package main

import (
	"io"
	"net/http"
)

func download(url string) ([]byte, error) {
	// use http like a client to request a URL from
	// a server and print the response of what it gets
	if rsp, err := http.Get(url); err == nil {
		if bts, err := io.ReadAll(rsp.Body); err == nil {
			// combining "if" with "err == nil" is more concise
			// but rsp, bts, err are only in scope inside the
			// block, not outside block
			return bts, nil
		}
	}
	return []byte{}, nil
}

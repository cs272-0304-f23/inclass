package main

import (
	"io"
	"net/http"
)

func download(url string) ([]byte, error) {
	if rsp, err := http.Get(url); err == nil {
		if b, err := io.ReadAll(rsp.Body); err == nil {
			// the scope of rsp, b, err are ONLY inside
			// the "if" clause, not outside
			return b, nil
		}
	}
	return []byte{}, nil
}

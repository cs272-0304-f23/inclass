package main

import (
	"fmt"
)

func main() {

	i := 0
	f := func() {
		func() {
			i = i + 1
			fmt.Println(i)
		}()
	}
	f()
	f()

	// sketch of using tc.text from parent scope in the anon
	// function which does the http handler
	for idx, tc := range TestCase {
	svr := httptest.NewServer(handler(func(){rw, r}(w.Write(tc.text)))
	}
}

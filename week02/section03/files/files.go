package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

// ./flags -input_file=foo.txt

func main() {
	var sptr *string
	sptr = flag.String("input_file", "", "name of the file to read")
	flag.Parse()

	f, err := os.Open(*sptr)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	s, err := io.ReadAll(f)
	if err == nil {
		fmt.Println(string(s))
	}
}

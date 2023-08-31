package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

// foo -input_file=foo.txt -count=99

func main() {
	var stringPtr *string
	stringPtr = flag.String("input_file", "", "file to read")
	// Fills in ptr values with parsed args from shell
	flag.Parse()

	f, err := os.Open(*stringPtr)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	if b, err := io.ReadAll(f); err == nil {
		fmt.Println(string(b))
	}

}

package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

// foo -input_file=foo.txt -count=99

func main() {
	var stringPtr *string
	stringPtr = flag.String("input_file", "", "file to read")
	// Fills in ptr values with parsed args from shell
	flag.Parse()

	lf, err := os.Create("log.txt")
	if err == nil {
		log.SetOutput(lf)
	} else {
		fmt.Println(err)
	}
	defer lf.Close()

	f, err := os.Open(*stringPtr)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	if b, err := io.ReadAll(f); err == nil {
		log.Println(string(b))
	}

}

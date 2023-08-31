package main

import (
	"flag"
	"fmt"
)

// foo -input_file=foo.txt -count=99

func main() {
	var stringPtr *string
	stringPtr = flag.String("input_file", "", "file to read")

	intPtr := flag.Int("count", 0, "count of things")

	// Fills in ptr values with parsed args from shell
	flag.Parse()

	fmt.Println("input_file: ", *stringPtr)
	fmt.Println("count: ", *intPtr)
}

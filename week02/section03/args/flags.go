package main

import (
	"flag"
	"fmt"
)

// ./flags -input_file=foo.txt

func main() {
	var sptr *string
	sptr = flag.String("input_file", "", "name of the file to read")
	flag.Parse()

	fmt.Println("sptr", *sptr)
}

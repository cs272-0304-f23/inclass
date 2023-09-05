package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	var sptr *string
	sptr = flag.String("input_file", "", "name of the file to read")
	flag.Parse()

	lf, err := os.Create("log.txt")
	if err == nil {
		log.SetOutput(lf)
	} else {
		fmt.Println(err)
	}

	f, err := os.Open(*sptr)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	s, err := io.ReadAll(f)
	if err == nil {
		log.Println(string(s))
	}
}

package main

import (
	"flag"
	"fmt"
)

func main() {
	// -download=https://www.usfca.edu
	urlPtr := flag.String("download", "", "URL to download")
	flag.Parse()
	body, err := download(*urlPtr)
	if err == nil {
		fmt.Println(string(body))
	}
}

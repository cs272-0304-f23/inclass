package main

import (
	"flag"
	"fmt"
)

func main() {
	urlPtr := flag.String("download", "", "URL to download")
	flag.Parse()
	body, err := download(*urlPtr)
	if err != nil {
		fmt.Println(body)
	}
}

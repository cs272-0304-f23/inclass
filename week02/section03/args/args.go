package main

import (
	"fmt"
	"os"
)

func main() {
	for idx, a := range os.Args {
		fmt.Printf("Args[%d] = %s\n", idx, a)
	}
}

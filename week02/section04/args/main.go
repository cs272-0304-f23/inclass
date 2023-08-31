package main

import (
	"fmt"
	"os"
)

func main() {
	for idx, a := range os.Args {
		fmt.Printf("args[%d] = %s\n", idx, a)
	}
}

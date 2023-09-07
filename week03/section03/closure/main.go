package main

import (
	"fmt"
)

func main() {
	i := 0

	f := func() {
		i = i + 1
	}

	f()
	f()

	fmt.Println("i: ", i)
}

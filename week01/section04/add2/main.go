package main

import (
	"fmt"
)

func add2(a, b int) int {
	return a + b
}

func main() {
	i := add2(1, 2)
	fmt.Println("i: ", i)
}

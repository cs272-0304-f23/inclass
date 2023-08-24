package main

import "fmt"

func main() {
	// Variables are zero-initialized in go
	var i int
	fmt.Println("i: ", i)

	// s is an empty string
	var s string
	fmt.Println("s: ", s)
}

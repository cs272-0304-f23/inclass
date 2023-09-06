package main

import (
	"fmt"
)

func main() {
	// anonymous functions can be defined and called within another function
	func(s string) {
		fmt.Println(s)
	}("hello there")

	f := func(s string) {
		fmt.Println(s)
	}

	// functions can be assigned to a variable
	f("hiya")
}

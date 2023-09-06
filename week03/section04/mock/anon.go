package main

import (
	"fmt"
)

func f(s string) {
	fmt.Println(s)
}

func main() {

	// anonymous function called inline using (args)
	func(s string) {
		fmt.Println(s)
	}("hello")

	// anonymous function assigned to a variable, and called
	f := func(s string) {
		fmt.Println(s)
	}

	f("hiya")
}

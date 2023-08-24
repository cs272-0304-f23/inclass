package main

import "fmt"

// Add2 is a function. It takes two integer args and
// returns an int
func Add2(a, b int) int {
	return a + b
}

func main() {
	i := Add2(1, 2)
	fmt.Println("i: ", i)
}

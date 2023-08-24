package main

import "fmt"

func main() {
	var i int = 1
	fmt.Println("i: ", i)

	// Short declaration: compiler looks at RHS to get type
	// of variable on LHS. Other examples:
	// s := "foobar"
	// obj := foobar()

	j := 2
	fmt.Println("j: ", j)

	fmt.Printf("i: %d j: %d\n", i, j)
}

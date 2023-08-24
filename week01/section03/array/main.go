package main

import "fmt"

func main() {
	// Declare an array of 3 integers
	// In go, array length is immutable, so arrays are
	// not that useful
	var arr [3]int

	// Initialize one value
	arr[0] = 2

	// arr[1] and arr[2] are zero due to zero-initialization
	for idx, val := range arr {
		fmt.Printf("arr[%d] = %d\n", idx, val)
	}
}

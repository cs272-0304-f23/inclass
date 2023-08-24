package main

import "fmt"

func PrintSlice(name string, s []int) {
	// range iterates the given slice, returning the index and value
	// for each element of the slice
	for idx, val := range s {
		fmt.Printf("%s[%d] = %d\n", name, idx, val)
	}
}

func main() {
	// Use short decl to declare a slice of ints using []
	// and initialize it to empty using {}
	s := []int{}

	// Append a couple elements to the slice
	s = append(s, 1)
	s = append(s, 2)

	// Append is "variadic" so we can pass multiple elements
	s = append(s, 3, 4, 5)
	PrintSlice("s", s)

	// Get a subset of slice elements: 0 is min, 3 is max
	// Returns elements at indexes 0, 1, 2
	s2 := s[0:3]
	s2 = append(s2, s[len(s)-1])
	PrintSlice("s2", s2)

	// Append a slice to another slice using ...
	s3 := []int{6, 7}
	s4 := append(s2, s3...)
	PrintSlice("s4", s4)
}

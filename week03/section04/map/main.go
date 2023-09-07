package main

import (
	"fmt"
)

func main() {
	// key is a string, value is an int
	// m["foo"] = 1
	// var m map[string]int

	// make() both allocates and initializes (to zero)
	// a new hash map
	m := make(map[string]int)

	m["foo"] = 1
	m["bar"] = 2

	fmt.Printf("m[foo] = %d\n", m["foo"])

	// to iterate the values of a hash table
	for key, val := range m {
		fmt.Printf("m[%s] = %d\n", key, val)
	}

	// to check if a hash table contains a key
	if _, ok := m["baz"]; ok {
		fmt.Println("m contains baz")
	} else {
		fmt.Println("m does not contain baz")
	}

	// the value for each key is unique
	// here, the 1 is gone, replace by the 3
	m["foo"] = 3
	fmt.Printf("m[foo] = %d\n", m["foo"])

	delete(m, "foo")
	if _, ok := m["foo"]; ok {
		fmt.Println("m contains foo")
	} else {
		fmt.Println("m does not contain foo")
	}

}

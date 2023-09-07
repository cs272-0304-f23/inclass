package main

import (
	"fmt"
)

func main() {
	// use make to allocate and initialize the map to zero
	m := make(map[string]int)

	m["foo"] = 1
	fmt.Printf("m[foo] = %d\n", m["foo"])

	// since "foo" hashes to the same bucket each time,
	// the prior value is replaced by a new one
	m["foo"] = 2
	fmt.Printf("m[foo] = %d\n", m["foo"])

	// we can range over the map, which gives key, value
	// (not idx, val like ranging over a slice)
	m["bar"] = 3
	for key, val := range m {
		fmt.Printf("m[%s] = %d\n", key, val)
	}

	// to check if a key exists in a map
	if val, ok := m["baz"]; ok {
		fmt.Println("exists: ", val)
	} else {
		fmt.Println("does not exist")
	}
}

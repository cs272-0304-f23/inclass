package main

import "fmt"

type Customer struct {
	Name    string
	Balance float32
}

func printSlice(s []*Customer) {
	for idx, pc := range s {
		fmt.Printf("s[%d] = %v", idx, *pc)
	}
}
func main() {
	s := []*Customer{}
	// s := make([]*Customer, 0)

	s = append(s, &Customer{"Noah", 2.0})
	printSlice(s)
}

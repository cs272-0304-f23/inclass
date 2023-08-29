package main

import "fmt"

type Customer struct {
	Name    string
	Balance float32
}

func makeCustomer(name string, balance float32) *Customer {
	// Stack escape: return pointer to a local variable
	// and Go compiler allocates it on the heap

	c := Customer{
		Name:    name,
		Balance: balance,
	}
	return &c
}

func main() {
	pc := makeCustomer("Phil", 0.0)

	s := make([]*Customer, 0)
	s = append(s, pc)

	// Creates customer using struct literal with
	// values in the same order as declared (Name, Balance)
	s = append(s, &Customer{"Logan", 1.0})

	for idx, p := range s {
		fmt.Printf("p[%d]: %v\n", idx, *p)
	}

	// some more code which doesn't reference s or pc
	// pc would be eligible for GC
}

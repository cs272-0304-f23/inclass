package main

import "fmt"

type Customer struct {
	Name    string
	Balance float32
}

func printCustomer(c *Customer) {
	fmt.Printf("%p", c)
}

// c is a "stack escape" where we return the address
// to the caller. Go puts the memory for "c" on the heap
// waiting to get freed by
func makeCustomer(name string, balance float32) *Customer {
	return &Customer{
		Name:    name,
		Balance: balance,
	}
}

func main() {
	pc := makeCustomer("Phil", 0.0)
	printCustomer(pc)

	// code which doesn't access the memory from pc
	//....

}

package main

import "fmt"

// Declare a new type of name Customer
type Customer struct {
	Name    string
	Balance float32
}

func PrintCustomer(c *Customer) {
	// In go, use '.' to dereference, not '->' like in C
	fmt.Println("Name: ", c.Name)
	fmt.Println("Balance: ", c.Balance)
}

func main() {
	// Use C-like approach to declare and initialize
	var c Customer
	c.Name = "Phil"
	c.Balance = 1.0
	PrintCustomer(&c)

	// Use short decl to declare and initialize
	// as a struct literal
	c2 := Customer{
		Name:    "Logan",
		Balance: 2.0,
	}
	PrintCustomer(&c2)

	// Create a struct literal in a func arg list
	PrintCustomer(&Customer{
		Name:    "Su",
		Balance: 3.0,
	})
}

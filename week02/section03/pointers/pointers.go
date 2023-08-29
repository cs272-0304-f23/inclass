package main

import "fmt"

type Customer struct {
	Name    string
	Balance float32
}

func printCustomer(c *Customer) {
	fmt.Printf("%v", *c)
}

func main() {
	c := Customer{
		Name:    "Phil",
		Balance: 0.0,
	}
	printCustomer(&c)

	var c2 *Customer = new(Customer)
	c2.Name = "David"
	c2.Balance = 1.0
	printCustomer(c2)
}

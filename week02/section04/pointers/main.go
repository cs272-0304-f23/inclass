package main

import "fmt"

type Customer struct {
	Name    string
	Balance float32
}

func main() {
	var pc *Customer = new(Customer)
	pc.Name = "Phil"
	pc.Balance = 0.0

	fmt.Printf("pc: %v", *pc)
}

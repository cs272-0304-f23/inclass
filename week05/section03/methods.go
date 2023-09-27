package main

import (
	"fmt"
)

// Interfaces are like abstract base classes
type Eater interface {
	Eat(food string)
}

type Dog struct {
	Name string
	Age  int
	Food string
}

// Dog implements the Eater interface because
// it implements all the methods of the interface
// matching types of arguments and return value
func (d *Dog) Eat(food string) {
	d.Food = food
}

func (d Dog) Print() {
	fmt.Println(d)
}

// If you don't use a pointer, you get the struct by value
// so Go makes another copy, we increment Age in the copy
// and then the copy is destroyed

// If you want changes in a method to be reflected, you
// have to use a pointer in the receiver, rather than value
func (d *Dog) Birthday() {
	d.Age++
	d.Print()
}

func main() {
	d := Dog{"Goldy", 11, ""}
	d.Print()
	d.Birthday()
	d.Print()
	d.Eat("kibble")
	d.Print()
}

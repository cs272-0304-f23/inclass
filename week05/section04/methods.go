package main

import (
	"fmt"
)

// C++ example; Java similar
// class ABC { void m() = 0; }
// class Derived : public ABC { void m(); }

type Eater interface {
	Eat(food string)
}

type Dog struct {
	Name string
	Age  int
	Ate  string
}

func (d *Dog) Eat(food string) {
	d.Ate = food
}

// The receiver d *Dog is passed by reference
// If the method mutates the object, use this
func (d *Dog) Birthday() {
	d.Age++
}

// The receiver d Dog is passed by value
// A copy of the struct is made
func (d Dog) Print() {
	fmt.Println(d)
}

func main() {
	g := Dog{"Goldy", 11, ""}
	g.Print()
	g.Birthday()
	g.Print()
	g.Eat("kibble")
	g.Print()
}

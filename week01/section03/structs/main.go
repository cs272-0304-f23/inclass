package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func PrintPerson(p *Person) {
	// Go uses '.' to dereference a ptr, just the same
	// as if the had the instance in a local variable (stack)
	fmt.Println("Name: ", p.Name)
	fmt.Println("Age: ", p.Age)
}

func main() {
	// Declare an instance of a Person
	var p Person
	p.Name = "David"
	p.Age = 20

	PrintPerson(&p)

	// Another using short decl and a "struct literal"
	p2 := Person{
		Name: "Phil",
		Age:  100,
	}
	PrintPerson(&p2)

	PrintPerson(&Person{
		Name: "Kerem",
		Age:  21,
	})

	PrintPerson(nil)
}

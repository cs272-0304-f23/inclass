package main

import "fmt"

func main() {
	// Declare s as a string and assign the string literal
	s := "foobar"
	fmt.Println("s:", s)

	// Assign a UTF-8 string literal to s2
	var s2 string = "føøbår"
	fmt.Println("s2: ", s2)

	// Use range to iterate the UTF-8 encoded string
	// returning the index of each char (or rune)
	for idx, val := range s2 {
		fmt.Printf("s2[%d] = %#U\n", idx, val)
	}

	for _, val := range s2 {
		fmt.Printf("s2 char = %#U\n", val)
	}
}

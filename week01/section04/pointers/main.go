package main

import "fmt"

func Add1(pi *int) {
	*pi = *pi + 1
}

func main() {
	j := 1
	i := &j
	Add1(i)
	fmt.Println("j: ", j)
}

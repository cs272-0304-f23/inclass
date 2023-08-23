package main

import "fmt"

func main() {
	var s string = "føøbår"
	fmt.Println("s: ", s)

	for _, r := range s {
		fmt.Printf("%#U", r)
	}

	// for i := 0; i < len(s); i++ {
	// fmt.Printf("s[%d] = %c\n", i, s[i])
	// }
}

package main

import (
	"fmt"
	"time"
)

func iterate(how string, ntimes int) {
	for i := 0; i < ntimes; i++ {
		// Putting in a Sleep here makes the goroutine not ready
		time.Sleep(10 * time.Millisecond)
		fmt.Println(how, i)
	}
}

func main() {

	// Synchronous operation. Main goroutine waits until
	// iterate is done before running the Println
	iterate("synchronous", 10)

	go iterate("goroutine1", 10)

	go iterate("goroutine2", 10)

	// A closure can also be a goroutine, accessing
	// variables from its parent scope in the parent goroutine
	i := 1
	go func(s string) {
		fmt.Println(s, i)
	}("literal")

	// Sleep is an ugly hack to wait "long enough" for
	// the goroutine to get scheduled and execute
	time.Sleep(1 * time.Second)

	fmt.Println("done")
}

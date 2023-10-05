package main

import (
	"fmt"
	"os"
	"strconv"
)

// ch is a channel that integers go into
func generate(ch chan<- int) {
	// infinite loop - no terminating condition
	for i := 0; ; i++ {
		ch <- i
	}
}

func main() {
	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("must be a number")
		return
	}

	ch := make(chan int)
	go generate(ch)

	for {
		i := <-ch
		fmt.Println(i)
		if i == num {
			break
		}
	}
}

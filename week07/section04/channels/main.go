package main

import (
	"fmt"
	"os"
	"strconv"
)

// A channel is a concurrency-safe way to communicate between
// goroutines

// chan <- int means that we will put integers into this channel
func generate(ch chan<- int) {
	for i := 0; ; i++ {
		ch <- i // put i into this channel
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
		// read from the channel, when ready, into i
		// if ch is not ready, <-ch blocks
		i := <-ch
		fmt.Println(i)
		if i == num {
			break
		}
	}

}

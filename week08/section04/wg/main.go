package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	const MAX_WG = 10
	const MAX_WORK = 100

	// unbuffered channel of integers
	ch := make(chan int)

	for i := 0; i < MAX_WG; i++ {
		go func(i int) {
			// add this goroutine to the wg
			wg.Add(1)

			// range over channel until closed
			for j := range ch {
				// simulated work
				time.Sleep(100 * time.Millisecond)
				fmt.Printf("%d %d\n", i, j)
			}

			// tell the wg this goroutine is done
			wg.Done()
		}(i)

		for i := 0; i < MAX_WORK; i++ {
			// write i into the channel
			ch <- i
		}
	}
	// stop ranging over the channel
	close(ch)

	// wait for all of the goroutines to finish
	wg.Wait()
}

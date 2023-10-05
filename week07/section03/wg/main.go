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

	ch := make(chan int)

	for i := 0; i < MAX_WG; i++ {
		go func(i int) {
			// Add this goroutine to the wg
			wg.Add(1)

			// Busy work for this example. A real example
			// would do real work here
			for j := range ch {
				time.Sleep(100 * time.Millisecond)
				fmt.Printf("%d,%d\n", i, j)
			}
			// Remove this goroutine from the wg
			wg.Done()
		}(i)
	}

	for i := 0; i < MAX_WORK; i++ {
		// put each integer into the channel
		ch <- i
	}
	// closing the channel makes range stop
	close(ch)

	// Wait blocks until all worker goroutines
	// have called wg.Done()
	wg.Wait()
	fmt.Println("Wait is done")
}

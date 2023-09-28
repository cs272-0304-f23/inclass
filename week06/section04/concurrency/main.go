package main

import (
	"fmt"
	"time"
)

func doStuff(typ string, howMany int) {
	for i := 0; i < howMany; i++ {
		if i%1000 == 0 {
			fmt.Println(typ, i)
			// Introducing sleep here makes the routine
			// not ready to execute, simulating disk/network/memory
			// latency. While waiting, another goroutine will be scheduled
			time.Sleep(10 * time.Millisecond)
		}
	}
}

// Code in main, or called synchronously from main
// runs in the "main" goroutine
func main() {

	// When calling a function synchronously, we wait until
	// it returns to do the next things. "sequential execution"
	doStuff("main", 10000)

	// If we call a function as a goroutine, we don't wait for it
	// Here, doStuff runs in its own goroutine, scheduled by
	// the Go runtime scheduler
	go doStuff("goroutine1", 10000)

	// Goroutine scheduling order is not guaranteed
	go doStuff("goroutine2", 10000)

	// Using time.Sleep is a grungy hack but makes the point
	// about async scheduling in this example.
	time.Sleep(1 * time.Second)

	fmt.Println("done!")
}

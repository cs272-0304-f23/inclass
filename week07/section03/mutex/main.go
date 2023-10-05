package main

import (
	"fmt"
	"sync"
	"time"
)

type Index struct {
	m  map[string]int
	mu sync.Mutex // mutual exclusion
}

func (i *Index) Increment(s string) {
	// Acquire the lock on the mutex for this this struct
	// to anyone else wanting to mutate has to wait for
	// our lock to be released
	i.mu.Lock()
	defer i.mu.Unlock()

	i.m[s]++
}

func (i *Index) Decrement(s string) {
	i.mu.Lock()
	defer i.mu.Unlock()

	i.m[s]--
}

func main() {
	i := &Index{
		m: make(map[string]int),
	}

	for {
		go i.Increment("foo")
		go i.Decrement("foo")
		fmt.Println(i.m["foo"])
		time.Sleep(100 * time.Millisecond)
	}
}

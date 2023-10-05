package main

import (
	"sync"
	"time"
)

type Index struct {
	m  map[string]int
	mu sync.Mutex
}

func NewIndex() *Index {
	return &Index{
		m: make(map[string]int),
	}
}

func (i *Index) Increment(s string) {
	// Acquire a lock on the index so no other goroutines
	// can change it out from under us
	i.mu.Lock()
	defer i.mu.Unlock()

	i.m[s]++
}

func (i *Index) Decrement(s string) {
	// Acquire a lock on the index so no other goroutines
	// can change it out from under us
	i.mu.Lock()
	defer i.mu.Unlock()

	i.m[s]--
}

func main() {
	i := NewIndex()

	for {
		go i.Increment("foo")
		go i.Decrement("foo")
		time.Sleep(100 * time.Millisecond)
	}
}

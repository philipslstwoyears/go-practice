package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	count int
	mu    sync.Mutex
}

func (c *Counter) Increment() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

func NewCounter() *Counter {
	return &Counter{}
}

func main() {
	counter := NewCounter()
	wg := sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			counter.Increment()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(counter.count)
}

package main

import (
	"fmt"
	"sync"
)

func main() {
	m := make(map[string]int)
	var keys = []string{"sfldsfd", "dfaospdjf", "adoigao", "iotp4upiow"}
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	for value, key := range keys {
		wg.Add(1)
		go func() {
			mu.Lock()
			m[key] = value
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(m)
}

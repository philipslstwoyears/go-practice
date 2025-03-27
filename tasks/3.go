package main

import (
	"fmt"
	"sync"
)

//Дана последовательность чисел: 2,4,6,8,10.
//Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.

func main() {
	m := []int{2, 4, 6, 8, 10}
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	var result int
	for _, num := range m {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			result += num * num
			mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println(result)
}

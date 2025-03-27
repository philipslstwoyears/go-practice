package main

import (
	"fmt"
	"sync"
)

// Написать программу,
// которая конкурентно рассчитает значение квадратов чисел взятых из массива
// (2,4,6,8,10) и выведет их квадраты в stdout.
func main() {
	wg := &sync.WaitGroup{}
	mass := []int{2, 4, 6, 8, 10}
	for _, num := range mass {
		wg.Add(1)
		go func() {
			fmt.Println(num * num)
			wg.Done()
		}()
	}
	wg.Wait()
}

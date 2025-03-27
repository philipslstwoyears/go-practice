package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	slnums := []int{
		1, 4, 2, 57, 59, 5, 3,
	}
	input := make(chan int)
	output := make(chan int)
	go func() {
		for num := range input {
			output <- num * 2

		}
		close(output)
	}()
	go func() {
		for _, slnum := range slnums {
			input <- slnum
		}
		close(input)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for num := range output {
			fmt.Println(num)
		}
	}()
	wg.Wait()
}

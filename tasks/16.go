package main

import (
	"log"
	"sort"
)

func main() {
	ms := []int{
		1, 2, 5, 4, 5, 6, 7, 124, 9, 1, 11, 2, 13, 14, 5, 16, 17, 18, 19,
	}
	sort.Ints(ms)
	log.Println(ms)
}

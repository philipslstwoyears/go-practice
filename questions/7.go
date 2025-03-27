package main

import "fmt"

func main() {
	m := make(map[int]int)
	m[0] = 1
	m[2] = 281
	m[4] = 281
	m[3] = 281
	m[1] = 124
	for key, value := range m {
		fmt.Println(key, value)
	}
}

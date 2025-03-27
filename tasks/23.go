package main

import "fmt"

func main() {
	n := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var i int
	fmt.Scan(&i)
	n = append(n[:i], n[i+1:]...)
	fmt.Println(n)
}

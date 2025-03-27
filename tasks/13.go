package main

import "fmt"

func main() {
	a := 64
	b := 102
	fmt.Printf("Первое число:%d, второе число:%d\n", a, b)
	//a, b = b, a
	var w int
	w = a
	a = b
	b = w
	fmt.Printf("Первое число:%d, второе число:%d\n", a, b)
}

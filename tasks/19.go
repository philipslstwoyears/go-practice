package main

import (
	"fmt"
)

func main() {
	var str string
	fmt.Scan(&str)
	n := []rune(str)
	var res []rune
	for i := len(n) - 1; i >= 0; i-- {
		res = append(res, n[i])
	}
	fmt.Println(string(res))
}

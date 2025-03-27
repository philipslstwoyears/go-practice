package main

import (
	"fmt"
	"strings"
)

func rev(str string) string {
	words := strings.Fields(str)
	res := make([]string, len(words))
	for i := 0; i < len(words); i++ {
		res[i] = words[len(words)-i-1]
	}
	return strings.Join(res, " ")
}

func main() {
	fmt.Println(rev("Hello World"))
}

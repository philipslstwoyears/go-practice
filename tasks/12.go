package main

import "fmt"

func main() {
	sl := []string{"dog", "dog", "cat"}
	res := make(map[string]struct{})
	for _, s := range sl {
		res[s] = struct{}{}
	}
	fmt.Println(res)
}

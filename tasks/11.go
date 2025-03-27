package main

import "fmt"

func main() {
	set1 := map[int]struct{}{
		1: {},
		2: {},
		3: {},
		4: {},
	}
	set2 := map[int]struct{}{
		3: {},
		4: {},
		5: {},
		6: {},
		7: {},
	}
	fmt.Println(intersection(set1, set2))
}
func intersection(a, b map[int]struct{}) map[int]struct{} {
	res := make(map[int]struct{})
	for key := range a {
		if _, ok := b[key]; ok {
			res[key] = struct{}{}
		}
	}
	return res
}

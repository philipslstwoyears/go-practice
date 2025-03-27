// b b a

package main

import "fmt"

func main() {
	slice := []string{"a", "a", "v"}

	func(slice []string) {
		slice = append(slice, "a")
		slice[0] = "b"
		slice[1] = "b"
		fmt.Print(slice) // [b,b,v,a]
	}(slice)
	fmt.Print(slice) // [a,a,v]
}

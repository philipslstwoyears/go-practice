package main

import "fmt"

type Human struct {
	name string
	age  int
}

func (h *Human) Speak() {
	fmt.Printf("Hello my name is %s", h.name)
}

type Action struct {
	Human
}

func main() {
	action := Action{
		Human{
			name: "L",
			age:  12,
		},
	}
	action.Speak()
}

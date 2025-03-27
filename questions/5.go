// struct{}{} занимает 0 байт потому что эта структура не содержит полей
// хоть каждый struct занимает 0 байт все равно у каждой такой структуры разные адреса

package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var s struct{}
	fmt.Println(unsafe.Sizeof(s))
}

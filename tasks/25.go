package main

import (
	"fmt"
	"time"
)

func sleep(t uint) {
	<-time.After(time.Duration(t) * time.Second)
}

func main() {
	var t uint
	fmt.Scan(&t)
	for i := 0; i < int(t); i++ {
		fmt.Println("TICK")
		sleep(1)
	}
}

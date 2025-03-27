package main

import "time"

// Разработать программу, которая будет последовательно отправлять значения в канал,
//а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.

func main() {
	st := new(struct{})
	ch := make(chan struct{})
	go func() {
		time.Sleep(1 * time.Second)
		ch <- *st
	}()
	<-ch
}

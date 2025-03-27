// при использовании обычного мьютекса есть эксклюзивный доступ только у одной горутины

// при использовании RW мьютекса мы можем читать из разных горутин одновременно
// но записывать записывает "эксклюзивная" горутина
package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func main() {
	var counter int
	var rw sync.RWMutex
	wg := sync.WaitGroup{}

	// читатель
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			rw.RLock()
			log.Println("Пользователь с id:", id, "читает")
			time.Sleep(time.Second * 2)
			rw.RUnlock()
			wg.Done()
		}(i)
	}

	// писатель
	wg.Add(1)
	go func() {
		defer wg.Done()
		rw.Lock()
		log.Println("Писатель начал")
		counter++
		time.Sleep(time.Second * 4)
		log.Println("Писатель закончил")
		rw.Unlock()
	}()
	wg.Wait()
	fmt.Printf("Итог: %d", counter)
}

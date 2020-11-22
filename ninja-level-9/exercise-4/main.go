package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var m sync.Mutex
	increase := 0
	gr := 100

	wg.Add(gr)

	for i := 0; i < gr; i++ {
		go func() {
			m.Lock()
			value := increase
			value++
			increase = value
			fmt.Println(increase)
			m.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Final value of increase:", increase)
}

package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	increase := 0
	gr := 100

	wg.Add(gr)

	for i := 0; i < gr; i++ {
		go func() {
			value := increase
			runtime.Gosched()
			value++
			increase = value
			fmt.Println(increase)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Final value of increase:", increase)
}

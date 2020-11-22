package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup
	var increase int64
	gr := 100

	wg.Add(gr)

	for i := 0; i < gr; i++ {
		go func() {
			atomic.AddInt64(&increase, 1)
			fmt.Println(atomic.LoadInt64(&increase))
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Final value of increase:", increase)
}

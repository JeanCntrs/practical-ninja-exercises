package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Printf("Number of CPUs at the start: %v\n", runtime.NumCPU())
	fmt.Printf("Number of Goroutines at the start: %v\n", runtime.NumGoroutine())

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		fmt.Println("Hi from the first goroutine")
		wg.Done()
	}()
	go func() {
		fmt.Println("Hi from the second goroutine")
		wg.Done()
	}()

	fmt.Printf("Number of CPUs in the middle: %v\n", runtime.NumCPU())
	fmt.Printf("Number of Goroutines in the middle: %v\n", runtime.NumGoroutine())

	wg.Wait()
	fmt.Println("About to finish main")

	fmt.Printf("Number of CPUs at the end: %v\n", runtime.NumCPU())
	fmt.Printf("Number of Goroutines at the end: %v\n", runtime.NumGoroutine())
}

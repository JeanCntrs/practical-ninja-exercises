package main

import "fmt"

func main() {
	defer foo()
	fmt.Println("hello")
}

func foo() {
	defer func() {
		fmt.Println("anonymous deferred function ran")
	}()
	fmt.Println("foo deferred ran")
}

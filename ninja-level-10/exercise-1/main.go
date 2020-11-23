package main

import (
	"fmt"
)

func main() {
	// With buffer
	c := make(chan int, 1)

	c <- 42

	// Without buffer
	/* go func() {
		c <- 42
	}() */

	fmt.Println(<-c)
}

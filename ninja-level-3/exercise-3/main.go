package main

import (
	"fmt"
	"time"
)

func main() {
	yearOfBirth := 1993
	currentYear := time.Now().Year()

	for i := yearOfBirth; i <= currentYear; i++ {
		fmt.Println(i)
	}
}

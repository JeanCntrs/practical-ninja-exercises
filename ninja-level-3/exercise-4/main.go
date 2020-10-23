package main

import (
	"fmt"
	"time"
)

func main() {
	yearOfBirth := 1993
	currentYear := time.Now().Year()

	for {
		if yearOfBirth > currentYear {
			break
		}

		fmt.Println(yearOfBirth)
		yearOfBirth++
	}
}

package main

import (
	"fmt"
	"time"
)

func main() {
	yearOfBirth := 1993
	currentYear := time.Now().Year()

	for yearOfBirth <= currentYear {
		fmt.Println(yearOfBirth)
		yearOfBirth++
	}
}

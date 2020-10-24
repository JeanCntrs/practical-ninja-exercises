package main

import (
	"fmt"
	"strings"
)

func main() {
	x := "Ratman"
	if strings.ToLower(x) == "batman" {
		fmt.Println("Batman")
		return
	} else if strings.ToLower(x) == "robin" {
		fmt.Println("Robin")
	} else {
		fmt.Println("Nothing")
	}
}

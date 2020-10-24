package main

import (
	"fmt"
	"strings"
)

func main() {
	x := "Batman"
	if strings.ToLower(x) == "batman" {
		fmt.Println("Es igual")
		return
	}

	fmt.Println("No es igual")
}

package main

import "fmt"

func main() {
	x := [5]int{11, 22, 33, 44, 55}

	for index, value := range x {
		fmt.Println(index, value)
	}

	fmt.Printf("%T", x)
}

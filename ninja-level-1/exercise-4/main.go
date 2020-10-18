package main

import "fmt"

type number int

var x number

func main() {
	fmt.Println(x)
	fmt.Printf("El tipo de x es: %T\n", x)

	x = 42
	fmt.Println(x)
}

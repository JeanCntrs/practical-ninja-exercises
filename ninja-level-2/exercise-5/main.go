package main

import "fmt"

func main() {
	a := `Esto es un 
	string literal no 
		interpretado, "Hello World"`
	fmt.Println(a)
}

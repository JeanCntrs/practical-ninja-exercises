package main

import (
	"fmt"
	"udemy/Eduar-Tua/practical-ninja-exercises/ninja-level-12/exercise-1/dog"
)

type canine struct {
	name string
	age  int
}

func main() {
	fido := canine{
		name: "Fido",
		age:  dog.Age(10),
	}
	fmt.Println(fido)
	fmt.Println(fido)
}

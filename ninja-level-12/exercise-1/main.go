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
	c1 := canine{
		name: "Chester",
		age:  dog.Age(2),
	}
	fmt.Println(c1)
}

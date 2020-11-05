package main

import "fmt"

type person struct {
	name            string
	surname         string
	favoriteFlavors []string
}

func main() {
	p1 := person{
		name:    "Noah",
		surname: "Smith",
		favoriteFlavors: []string{
			"chocolate",
			"raspberry",
			"rum with raisins",
		},
	}

	p2 := person{
		name:    "Jack",
		surname: "Mason",
		favoriteFlavors: []string{
			"vainilla",
			"pineapple",
			"assortment",
		},
	}

	fmt.Println(p1.name)
	fmt.Println(p1.surname)
	for i, v := range p1.favoriteFlavors {
		fmt.Println("\t", i, v)
	}

	fmt.Println(p2.name)
	fmt.Println(p2.surname)
	for i, v := range p2.favoriteFlavors {
		fmt.Println("\t", i, v)
	}
}

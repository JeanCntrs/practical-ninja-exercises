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

	m := map[string]person{
		p1.surname: p1,
		p2.surname: p2,
	}

	for _, value := range m {
		fmt.Println(value.name)
		fmt.Println(value.surname)
		for i, flavorValue := range value.favoriteFlavors {
			fmt.Println(" ", i, flavorValue)
		}
		fmt.Println("---------")
	}
}

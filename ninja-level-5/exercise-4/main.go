package main

import "fmt"

func main() {
	anonymousStruct1 := struct {
		name    string
		friends map[string]int
		drinks  []string
	}{
		name: "James",
		friends: map[string]int{
			"Charlie": 123,
			"Thomas":  456,
			"Kyle":    789,
		},
		drinks: []string{"water", "vodka"},
	}

	fmt.Println(anonymousStruct1.name)
	fmt.Println(anonymousStruct1.friends)
	fmt.Println(anonymousStruct1.drinks)

	for key, value := range anonymousStruct1.friends {
		fmt.Println(key, value)
	}

	for index, value := range anonymousStruct1.drinks {
		fmt.Println(index, value)
	}

	fmt.Println("---------------------")

	anonymousStruct2 := []struct {
		name    string
		friends map[string]int
		drinks  []string
	}{
		{name: "William",
			friends: map[string]int{
				"David":   123,
				"Richard": 456,
				"Joseph":  789,
			},
			drinks: []string{"Absolut", "Baileys"}},
		{
			name: "Daniel",
			friends: map[string]int{
				"Damian": 123,
				"George": 456,
				"Jacob":  789,
			},
			drinks: []string{"Jägermeister", "Khlibniy"},
		},
	}

	for _, value := range anonymousStruct2 {
		fmt.Println(value.name)
		for key, value := range value.friends {
			fmt.Println("\t", key, value)
		}
		fmt.Println("")
		for index, value := range value.drinks {
			fmt.Println(index, value)
		}
		fmt.Println("")
	}
}

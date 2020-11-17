package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name    string   `json:"Nombre"`
	Surname string   `json:"Apellido"`
	Age     int      `json:"Edad"`
	Sayings []string `json:"Dichos"`
}

func main() {
	s := `[
		{"Nombre":"Eduar","Apellido":"Tua","Edad":32,"Dichos":["Cachicamo diciéndole a morrocoy conchudo","La mona, aunque se vista de seda, mona se queda","Poco a poco se anda lejos"]},
		{"Nombre":"Carlos","Apellido":"Pérez","Edad":27,"Dichos":["Camarón que se duerme se lo lleva la corriente","A ponerse las alpargatas que lo que viene es joropo","No gastes pólvora en zamuro"]},
		{"Nombre":"M","Apellido":"Hmmmm","Edad":54,"Dichos":["Ni lava ni presta la batea","Hijo de gato, caza ratón","Más vale pájaro en mano que cien volando"]}
	]`
	fmt.Println(s)

	var people []person
	err := json.Unmarshal([]byte(s), &people)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(people)

	for i, person := range people {
		fmt.Println("\tPersona #", i+1)
		fmt.Println("\t", person.Name, person.Surname, person.Age)
		for _, saying := range person.Sayings {
			fmt.Println("\t\t", saying)
		}
	}

}

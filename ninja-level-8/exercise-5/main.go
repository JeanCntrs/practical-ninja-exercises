package main

import (
	"fmt"
	"sort"
)

type usuario struct {
	Nombre   string
	Apellido string
	Edad     int
	Dichos   []string
}

type byAge []usuario

func (a byAge) Len() int           { return len(a) }
func (a byAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byAge) Less(i, j int) bool { return a[i].Edad < a[j].Edad }

type bySurname []usuario

func (a bySurname) Len() int           { return len(a) }
func (a bySurname) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a bySurname) Less(i, j int) bool { return a[i].Apellido < a[j].Apellido }

func main() {
	u1 := usuario{
		Nombre:   "Eduar",
		Apellido: "Tua",
		Edad:     32,
		Dichos: []string{
			"Cachicamo diciéndole a morrocoy conchudo",
			"La mona, aunque se vista de seda, mona se queda",
			"Poco a poco se anda lejos",
		},
	}

	u2 := usuario{
		Nombre:   "Carlos",
		Apellido: "Pérez",
		Edad:     27,
		Dichos: []string{
			"Camarón que se duerme se lo lleva la corriente",
			"A ponerse las alpargatas que lo que viene es joropo",
			"No gastes pólvora en zamuro",
		},
	}

	u3 := usuario{
		Nombre:   "Che",
		Apellido: "López",
		Edad:     54,
		Dichos: []string{
			"Ni lava ni presta la batea",
			"Hijo de gato, caza ratón",
			"Más vale pájaro en mano que cien volando",
		},
	}

	usuarios := []usuario{u1, u2, u3}
	for _, u := range usuarios {
		fmt.Println(u.Nombre, u.Apellido, u.Edad)
		sort.Strings(u.Dichos)
		for _, d := range u.Dichos {
			fmt.Println("\t", d)
		}
	}

	fmt.Println("------------")
	sort.Sort(byAge(usuarios))
	for _, u := range usuarios {
		fmt.Println(u.Nombre, u.Apellido, u.Edad)
		for _, d := range u.Dichos {
			fmt.Println("\t", d)
		}
	}

	fmt.Println("------------")
	sort.Sort(bySurname(usuarios))
	for _, u := range usuarios {
		fmt.Println(u.Nombre, u.Apellido, u.Edad)
		for _, d := range u.Dichos {
			fmt.Println("\t", d)
		}
	}
}

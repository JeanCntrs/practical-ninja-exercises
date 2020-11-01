package main

import "fmt"

func main() {
	x := map[string][]string{
		"eduar_tua":    {"computadoras", "montaña", "playa"},
		"carlos_ramon": {"leer", "comprar", "música"},
	}
	x["juan_bimba"] = []string{"helado", "pintar", "bailar"}
	fmt.Println(x)

	for key, value := range x {
		fmt.Println("Registro:", key)
		for k, v := range value {
			fmt.Println("\t", k, v)
		}
	}
}

package main

import "fmt"

func main() {
	xs1 := []string{"Batman", "Jefe", "Vestido de negro"}
	xs2 := []string{"Robin", "Ayudante", "Vestido de colores"}
	xss := [][]string{xs1, xs2}
	fmt.Println(xss)

	for i, reg := range xss {
		fmt.Println("Registro:", i)
		for j, val := range reg {
			fmt.Printf("\tÍndice de posición: %v\tvalor: %v\n", j, val)
		}
	}
}

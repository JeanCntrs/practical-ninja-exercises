package main

import "fmt"

type person struct {
	name string
}

type human interface {
	talk()
}

func (p *person) talk() {
	fmt.Println("Hi I´m", p.name)
}

func saySomething(h human) {
	h.talk()
}

func main() {
	p1 := person{name: "Juan"}

	saySomething(&p1)

	// El compilador usa direcciones para llamar a los métodos, p1 tiene dirección
	p1.talk()
}

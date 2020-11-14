package main

import "fmt"

type person struct {
	name    string
	surname string
	age     int
}

func (p person) introduceYourself() {
	fmt.Println("Hi, I´m", p.name, p.surname, "and have", p.age, "years old.")
}

func main() {
	p1 := person{
		name:    "James",
		surname: "Bond",
		age:     32,
	}

	p1.introduceYourself()
}

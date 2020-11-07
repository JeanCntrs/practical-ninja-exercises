package main

import "fmt"

type vehicle struct {
	doors  int
	colour string
}

type truck struct {
	vehicle
	fourWheels bool
}

type sedan struct {
	vehicle
	luxurious bool
}

func main() {
	c := truck{
		vehicle: vehicle{
			doors:  2,
			colour: "white",
		},
		fourWheels: true,
	}

	s := sedan{
		vehicle: vehicle{
			doors:  4,
			colour: "black",
		},
		luxurious: false,
	}

	fmt.Println(c)
	fmt.Println(s)
	fmt.Println(c.doors)
	fmt.Println(s.vehicle.doors)
}

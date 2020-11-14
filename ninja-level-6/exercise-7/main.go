package main

import "fmt"

var x int = 7
var g func() = func() {
	fmt.Println("g from outside")
}

func main() {
	f := func() {
		for i := 0; i <= 2; i++ {
			fmt.Println(i)
		}
	}

	f()
	fmt.Printf("%T\n", f)

	fmt.Println(x)
	fmt.Printf("%T\n", x)

	g()
	fmt.Printf("%T\n", g)

	fmt.Println("Ready!")

}

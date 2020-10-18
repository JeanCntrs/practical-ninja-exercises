package main

import "fmt"

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
	s := fmt.Sprintf("x=%v\ty=%v\tz=%v", x, y, z)
	fmt.Println(s)
}

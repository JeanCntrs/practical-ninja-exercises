package main

import "fmt"

func main() {
	si := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	f := foo(si...)
	fmt.Println(f)

	b := bar(si)
	fmt.Println(b)
}

func foo(si ...int) int {
	fmt.Println(si)
	total := 0
	for _, v := range si {
		total += v
	}
	return total
}

func bar(si []int) int {
	fmt.Println(si)
	total := 0
	for _, v := range si {
		total += v
	}
	return total
}

package main

import "fmt"

func main() {
	deporteFav := "surf"

	switch deporteFav {
	case "beisbol":
		fmt.Println("go to the stadium")
	case "hunt":
		fmt.Println("go to the jungle")
	case "surf":
		fmt.Println("go to the beach")
	}
}

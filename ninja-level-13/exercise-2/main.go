package main

import (
	"fmt"
	"udemy/Eduar-Tua/practical-ninja-exercises/ninja-level-13/exercise-2/quote"
	"udemy/Eduar-Tua/practical-ninja-exercises/ninja-level-13/exercise-2/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}

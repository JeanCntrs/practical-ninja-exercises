package main

import "fmt"

type customError struct {
	info string
}

func (ce customError) Error() string {
	return fmt.Sprintf("Here is the error: %v", ce.info)
}

func main() {
	ce1 := customError{info: "I need sleep more"}
	foo(ce1)
}

func foo(e error) {
	// assertion
	fmt.Println("foo ran -", e, "\n", e.(customError).info)
}

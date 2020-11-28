package word

import (
	"fmt"
	"testing"
	"udemy/Eduar-Tua/practical-ninja-exercises/ninja-level-13/exercise-2/quote"
)

func TestCount(t *testing.T) {
	c := Count(quote.SunAlso)
	if c != 1349 {
		t.Errorf("expected: %d got: %d", 1349, c)
	}
}

func TestUseCount(t *testing.T) {
	m := UseCount("one two three three three")
	for key, value := range m {
		switch key {
		case "one":
			if value != 1 {
				t.Errorf("expected: %d got: %d", 1, value)
			}
		case "two":
			if value != 1 {
				t.Errorf("expected: %d got: %d", 1, value)
			}
		case "three":
			if value != 3 {
				t.Errorf("expected: %d got: %d", 3, value)
			}
		}
	}
}

func ExampleCount() {
	fmt.Println(Count("one two three"))
	// Output: 3
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(quote.SunAlso)
	}
}

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount(quote.SunAlso)
	}
}

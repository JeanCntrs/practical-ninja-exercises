package dog

import (
	"fmt"
	"testing"
)

func TestAge(t *testing.T) {
	a := Age(10)
	if a != 70 {
		t.Errorf("Expected %d, Got %d", 70, a)
	}
}

func TestAgeTwo(t *testing.T) {
	a := AgeTwo(10)
	if a != 70 {
		t.Errorf("Expected %d, Got %d", 70, a)
	}
}

func ExampleAge() {
	fmt.Println(Age(10))
	// Output:
	// 70
}

func ExampleAgeTwo() {
	fmt.Println(AgeTwo(10))
	// Output: 70
}

func BenchmarkAge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Age(10)
	}
}

func BenchmarkAgeTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AgeTwo(10)
	}
}

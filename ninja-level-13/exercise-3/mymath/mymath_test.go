package mymath

import (
	"fmt"
	"testing"
)

func TestCenteredAvg(t *testing.T) {
	type test struct {
		data     []int
		response float64
	}

	tests := []test{
		{data: []int{1, 2, 3, 4, 5}, response: 3},
		{[]int{3, 5, 7, 8}, 6},
		{[]int{10, 20, 30, 40, 50}, 30},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 5},
	}

	for _, value := range tests {
		ca := CenteredAvg(value.data)
		if ca != value.response {
			t.Errorf("expected: %v got: %v", value.response, ca)
		}
	}
}

func ExampleCenteredAvg() {
	fmt.Println(CenteredAvg([]int{1, 2, 3, 4, 5}))
	// Output: 3
}

func BenchmarkCenteredAvg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CenteredAvg([]int{1, 2, 3, 4, 5, 6, 7, 1000000})
	}
}

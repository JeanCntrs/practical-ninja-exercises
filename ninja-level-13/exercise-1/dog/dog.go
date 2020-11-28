// Package dog allows us to better understand dogs
package dog

// Age gives us the age of a human in dog years
func Age(a int) int {
	return a * 7
}

// AgeTwo gives us the age of a human in dog years
func AgeTwo(a int) int {
	count := 0
	for i := 0; i < a; i++ {
		count += 7
	}
	return count
}

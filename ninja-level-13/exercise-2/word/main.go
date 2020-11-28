// Package word help us work with strings
package word

import "strings"

// UseCount returns us a map with de words and the number of times they appear
// no need to write an example for this one
// writing a test for this one is a bonus challenge; harder
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

// Count gives us the total number of words
func Count(s string) int {
	xs := strings.Fields(s)
	return len(xs)
}

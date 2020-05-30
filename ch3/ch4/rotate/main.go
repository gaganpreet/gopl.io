package main

import "fmt"

func rotate(s []int, idx int) []int {
	// Rotate an array s around idx
	r := s[idx:]
	for i := range s[:idx] {
		r = append(r, s[i])
	}
	return r
}

func main() {
	fmt.Println(rotate([]int{0, 1, 2, 3, 4, 5}, 2))
}

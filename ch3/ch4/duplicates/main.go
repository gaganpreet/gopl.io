// 4.5 Write an in-place fn to eliminate adjacent duplicates in a []string slice

package main

import "fmt"

// Eliminates (consecutive) duplicates
func eliminateDuplicates(strings []string) []string{
	i := 0
	var prev string
	for _, s := range(strings) {
		if s != prev {
			strings[i] = s
			i++
		}
		prev = s
	}
	return strings[:i]
}

func main() {
	r := eliminateDuplicates([]string{"hello", "world", "world", "hello", "hello", "world", "world"})
	fmt.Println(r)
}

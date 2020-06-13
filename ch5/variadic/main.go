package main

import "fmt"

func sum (vals ...int) int {
	result := 0
	for _, val := range(vals) {
		result += val
	}
	return result
}

func max (first int, vals ...int) (result int) {
	result = first
	for _, val := range(vals) {
		if val > result {
			result = val
		}
	}
	return
}

func min (first int, vals ...int) (result int) {
	for _, val := range(vals) {
		if val < result {
			result = val
		}
	}
	return
}

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5))
	fmt.Println(max(1, 2, 3, 4, 5))
	fmt.Println(min(1, 2, 3, 4, 5))
}

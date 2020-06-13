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

func Join (sep string, vals ...string) (result string) {
	currentSep := ""
	for i, val := range vals {
		result = fmt.Sprintf("%s%s%s", result, currentSep, val)
		if i == 0 {
			currentSep = sep
		}
	}
	return result
}

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5))
	fmt.Println(max(1, 2, 3, 4, 5))
	fmt.Println(min(1, 2, 3, 4, 5))

	stringArray := []string{"1", "2", "3"}
	fmt.Println(Join(" ", stringArray...))
}

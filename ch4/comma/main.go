package main

import "fmt"

func comma(s string) string {
	first, second := s[:len(s) - 3], s[len(s) - 3:]

	if len(first) < 3 {
		return s
	}
	return comma(first) + "," + second
}

func main() {
	fmt.Println(comma("12345"))
	fmt.Println(comma("134075971343401"))
}

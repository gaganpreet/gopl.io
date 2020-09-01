package main

import (
	"fmt"
	"sort"
)

func IsPalindrome(s sort.Interface) bool {
	l := s.Len()
	for i := 0; i < l / 2; i++ {
		j := l - i - 1 
		if !s.Less(i, j) && !s.Less(j, i) {
			continue
		}
		return false
	}
	return true
}

func main() {
	fmt.Println(IsPalindrome(sort.IntSlice([]int{1, 2, 3, 3, 2, 1})))
	fmt.Println(IsPalindrome(sort.IntSlice([]int{1, 2, 3, 2, 1})))
	fmt.Println(IsPalindrome(sort.IntSlice([]int{1, 2, 3, 5, 1})))
}

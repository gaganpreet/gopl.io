package main

import "fmt"

func IsAnagram(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	countsA := make(map[byte]int)
	countsB := make(map[byte]int)

	for i := range a {
		countsA[a[i]] += 1
		countsB[b[i]] += 1
	}

	for c := range countsA {
		if countsA[c] != countsB[c] {
			return false
		}
	}

	return true
}

func main() {

	fmt.Println(IsAnagram("abcd", "dbca"))
	fmt.Println(IsAnagram("aaaa", "dbca"))
	fmt.Println(IsAnagram("aaa", "dbca"))

}

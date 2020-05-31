// 4.6 Write an in-place function that squashes each run of adjacent Unicode spaces in a UTF-8-encoded [] byte slice into a single ASCII space.

package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func squashSpaces(s []byte) []byte {
	i := 0
	prevSpace := false
	skipped := 0
	for i < len(s)-skipped {
		r, size := utf8.DecodeRune(s[i:])
		if unicode.IsSpace(r) {
			if !prevSpace {
				s[i] = ' '
				skipped += size - 1
				copy(s[i + 1:], s[i + size :])
				i += 1
			} else {
				copy(s[i:], s[i+size:])
				skipped += size
			}
			prevSpace = true
		} else {
			prevSpace = false
			i += size
		}
	}
	fmt.Println()
	return s[:i]
}

func main() {
	s := "a        b   c         a    c"
	squashed := squashSpaces([]byte(s))
	fmt.Println(string(squashed))
}

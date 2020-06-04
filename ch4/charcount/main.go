package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

type RuneType string

const (
	RuneLetter RuneType = "letter"
	RuneNumber RuneType = "number"
	RuneSpace  RuneType = "space"
	RuneOther  RuneType = "other"
)

func main() {
	counts := make(map[rune]int)         // counts of Unicode characters
	typeCounts := make(map[RuneType]int) // counts of rune types
	var utflen [utf8.UTFMax + 1]int      // count of lengths of UTF-8 encodings
	invalid := 0

	in := bufio.NewReader(os.Stdin)

	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		var runeType RuneType
		if unicode.IsLetter(r) {
			runeType = RuneLetter
		} else if unicode.IsNumber(r) {
			runeType = RuneNumber
		} else if unicode.IsSpace(r) {
			runeType = RuneSpace
		} else {
			runeType = RuneOther
		}
		typeCounts[runeType] += 1
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
		}

		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	fmt.Print("\ntype\tcount\n")
	for i, n := range typeCounts {
		fmt.Printf("%s\t%d\n", i, n)
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}

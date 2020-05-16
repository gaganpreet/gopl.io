package main

import (
	"bytes"
	"fmt"
)

func comma(s string) string {
	first, second := s[:len(s) - 3], s[len(s) - 3:]

	if len(first) < 3 {
		return s
	}
	return comma(first) + "," + second
}

func commaBuffer(s string) string {
	// Exercise: non recursive comma using buffer
	var buf bytes.Buffer

	offset := len(s) % 3
	if offset == 0 {
		offset = 3
	}
	for ; len(s) > 0;  {
		buf.WriteString(s[:offset])
		s = s[offset:]
		if len(s) > 0 {
			buf.WriteString(",")
		}
		offset = 3
	}
	buf.WriteString(s)
	return buf.String()
}

func main() {
	fmt.Println(commaBuffer("12345"))
	fmt.Println(commaBuffer("134075971343401"))
}

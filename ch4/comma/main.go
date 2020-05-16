package main

import (
	"bytes"
	"fmt"
	"strings"
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

	if s[0] == '+' || s[0] == '-' {
		buf.WriteByte(s[0])
		s = s[1:]
	}

	suffix:= ""

	if decimalIndex := strings.Index(s, "."); decimalIndex != -1 {
		s, suffix = s[:decimalIndex], s[decimalIndex:]
	}

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
	buf.WriteString(suffix)
	return buf.String()
}

func main() {
	fmt.Println(commaBuffer("12345"))
	fmt.Println(commaBuffer("134075971343401"))
	fmt.Println(commaBuffer("+12345"))
	fmt.Println(commaBuffer("-12345"))
	fmt.Println(commaBuffer("12345.314341"))
}

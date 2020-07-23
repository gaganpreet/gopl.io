package main

import (
	"fmt"
	"io"
)

type StringReader struct {
	s string
	i int
}

func (sr *StringReader) Read (p []byte) (int, error) {
	if sr.i > len(sr.s) {
		return 0, nil
	}
	n := copy(p, sr.s[sr.i:])
	sr.i += n
	if sr.i >= len(sr.s) {
		return n, io.EOF
	}
	return n, nil
}

func NewReader (s string) *StringReader {
	sr := StringReader{s: s, i: 0}
	return &sr
}

func main() {
	r := NewReader("hello world as a test")
	b := make([]byte, 8)
	for {
		n, _ := r.Read(b)
		if n == 0 {
			break
		}
		fmt.Printf("Read %d bytes: <%s>\n", n, b[:n])
	}
}

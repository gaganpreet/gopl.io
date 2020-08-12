package main

import (
	"fmt"
	"io"
	"strings"
)

type LimitReaderStruct struct {
	r io.Reader
	limit int64
}

func (l *LimitReaderStruct) Read(p  []byte) (int, error) {
	if int64(len(p)) <= l.limit {
		n, err := l.r.Read(p)
		l.limit -= int64(n)
		return n, err
	} else {
		n, _ := l.r.Read(p[:l.limit])
		l.limit -= int64(n)
		return n, io.EOF
	}
}

func LimitReader(r io.Reader, n int64) LimitReaderStruct {
	return LimitReaderStruct{r: r, limit: n}
}

func main() {
	r := strings.NewReader("hello world")
	b := make([]byte, 5)

	limitReader := LimitReader(r, 20)
	for {
		n, _ := limitReader.Read(b)
		if n == 0 {
			break
		}
		fmt.Printf("Read %s bytes: <%s>\n", n, b[:n])
	}
}

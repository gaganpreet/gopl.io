package main

import (
	"bufio"
	"fmt"
)

type ByteCounter int
type WordCounter int
type LineCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) // convert int to ByteCounter
	return len(p), nil
}

func (c *WordCounter) Write(p []byte) (int, error) {
	count := 0
	var err error
	for {
		advance, token, _ := bufio.ScanWords(p, true)
		if len(token) == 0 {
			break
		}
		count++
		p = p[advance:]
	}
	*c += WordCounter(count)
	return count, err
}

func (c *LineCounter) Write(p []byte) (int, error) {
	count := 0
	var err error
	for {
		advance, token, _ := bufio.ScanLines(p, true)
		if len(token) == 0 {
			break
		}
		count++
		p = p[advance:]
	}
	*c += LineCounter(count)
	return count, err
}


func main() {
	var c WordCounter
	var l LineCounter
	c.Write([]byte("hello"))
	fmt.Println(c)

	c = 0
	var name = "Dolly"
	fmt.Fprintf(&c, "Hello, %s", name)
	fmt.Println(c)

	l.Write([]byte("hello"))
	fmt.Println(l)

	l = 0
	fmt.Fprintf(&l, "Hello, %s", name)
	fmt.Println(l)
}

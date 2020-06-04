package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("a"))
	c2 := sha256.Sum256([]byte("b"))

	fmt.Printf("%x, %x\n", c1, c2)
}

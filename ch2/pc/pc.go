package main

import (
	"fmt"
	"time"

	"gopl.io/ch2/popcount"
)

func main() {
	var i uint64

	start := time.Now()
	for i = 0; i < 1 << 32; i++ {
		popcount.PopCount(i)
	}
	fmt.Printf("With lookup: %v\n", time.Since(start).Seconds())

	start = time.Now()
	for i = 0; i < 1 << 32; i++ {
		popcount.PopCountLoop(i)
	}
	fmt.Printf("With for loop: %v\n", time.Since(start).Seconds())
}

package main

import "fmt"

func panicAndRecover() (x int) {
	defer func() {
		p := recover()
		if p != nil {
			x = 42
		}
	}()
	panic("failed: unrecoverable error")
}

func main() {
	fmt.Println(panicAndRecover())
}

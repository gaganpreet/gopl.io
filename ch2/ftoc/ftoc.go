package main

import "fmt"

func main() {
	freezing, boiling := 32.0, 212.0
	fmt.Printf("%fF is %fC\n", freezing, ftoc(freezing))
	fmt.Printf("%fF is %fC\n", boiling, ftoc(boiling))
}

func ftoc(temp float64) float64 {
	return 5 * (temp - 32) / 9;
}

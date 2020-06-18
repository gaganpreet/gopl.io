package main

import (
	"fmt"

	"gopl.io/ch6/intset"
)

func main() {
	var x, y intset.IntSet
	x.AddAll(1, 144, 9)
	fmt.Println(x.String())
	fmt.Println(x.Length())

	y.AddAll(9, 42)

	fmt.Println(x.Remove(144))
	fmt.Println(x.Remove(5))

	fmt.Println(y.String())
	fmt.Println(x.Length())
	
	x.UnionWith(&y)

	fmt.Println(&x)
	t := *x.Copy()
	fmt.Println(&x)
	fmt.Println(&t)

	x.Clear()
	fmt.Println(&x)
	x.Add(10)
	fmt.Println(&x)
}

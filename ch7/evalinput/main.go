// exercise 7.15 -> read an expression and parse it
package main

import (
	"bufio"
	"fmt"
	"os"

	"gopl.io/ch7/eval"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		expr, err := eval.Parse(input.Text())
		vars := make(map[eval.Var]bool)
		if err != nil {
			fmt.Errorf("error parsing expression: %s", err)
			continue
		}
		err = expr.Check(vars)
		if err != nil {
			fmt.Errorf("error parsing expression: %s", err)
			continue
		}
		env := make(map[eval.Var]float64)
		for varName := range vars {
			varVal := 0.0
			fmt.Printf("Value for %s: ", varName)
			fmt.Scanf("%f", &varVal)
			env[eval.Var(varName)] = varVal
		}
		fmt.Println(expr.Eval(env))
	}
}

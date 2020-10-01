package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"gopl.io/ch7/eval"
)

func evalRoute(w http.ResponseWriter, req *http.Request) {
	// eg:
	// http://localhost:8001/?expr=x-y&x=1&y
	// http://localhost:8001/?expr=pi*r*2&pi=3.14159&r=2
	exprString := req.URL.Query().Get("expr")
	expr, err := eval.Parse(exprString)
	if err != nil {
		fmt.Fprintf(w, "error parsing expression: %s", err)
		return
	}

	vars := make(map[eval.Var]bool)
	err = expr.Check(vars)
	if err != nil {
		fmt.Fprintf(w, "error parsing expression: %s", err)
		return
	}
	env := make(map[eval.Var]float64)

	for varName := range vars {
		val, err := strconv.ParseFloat(req.URL.Query().Get(string(varName)), 64)
		if err != nil {
			fmt.Fprintf(w, "missing value for %s in params", varName)
			return
		}
		env[eval.Var(varName)] = val
	}
	fmt.Fprintf(w, "Result of evaluating %s: %f", exprString, expr.Eval(env))
}

func main() {
	http.HandleFunc("/", evalRoute)
	log.Fatal(http.ListenAndServe(":8001", nil))
}

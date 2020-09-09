package expression

import (
	"fmt"
	"math"
	"testing"
)

func TestEval(t *testing.T) {
	tests := []struct {
		expr string
		env Env
		want string
	}{
		{"sqrt(A/pi)", Env{"A": 87616, "pi": math.Pi}, "167"},
		{"pow(x, 3) + pow(y, 3)", Env{"x": 12, "y":1}, "1729"},

	}
	var prevExpr string
	for _, test := range tests {
		// Print expr only when it changes
		if test.expr != prevExpr {
			fmt.Printf("\n%s\n", test.expr)
			prevExpr = test.expr
		}
		expr, err := Parse(test.expr)
		if err != nil {
			t.Error(err) // parse error
			continue
		}
		got := fmt.Sprintf("%.6g", expr.Eval(test.env))
		fmt.Printf("\t%v => %s\n", test.env, got)
		if got != test.want {
			t.Errorf("%s.Eval() in %v = %q, want %q\n",
		test.expr, test.env, got, test.want)
		}
	}
}

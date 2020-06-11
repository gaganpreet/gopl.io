package main

import (
	"fmt"
	"os"
)

type VisitedType int8

const (
	Unvisited VisitedType = iota
	TemporaryMark
	PermanentMark
)

var prereqs = map[string][]string{
	"algorithms":     {"data structures"},
	"calculus":       {"linear algebra"},
	// "linear algebra": {"calculus"}, // Exercise 5.11 to detect cycles
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	result, err := topoSort(prereqs)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	for i, course := range result {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) ([]string, error) {
	var order []string
	var err error
	state := make(map[string]VisitedType)
	var visitAll func()
	var visit func(start string)
	visit = func(start string) {
		if state[start] == TemporaryMark{
			err = fmt.Errorf("Found cycle for %s", start)
		}
		if state[start] == Unvisited {
			state[start] = TemporaryMark
			for _, item := range m[start] {
				visit(item)
			}
			order = append(order, start)
			state[start] = PermanentMark
		}
	}
	visitAll = func() {
		for item := range prereqs {
			if state[item] == Unvisited {
				visit(item)
			}
		}
	}
	visitAll()
	return order, err
}

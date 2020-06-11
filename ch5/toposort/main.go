package main

import (
	"fmt"
)

type VisitedType int8

const (
	Unvisited VisitedType = iota
	TemporaryMark
	PermanentMark
)

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
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
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) []string {
	var order []string
	state := make(map[string]VisitedType)
	var visitAll func()
	var visit func(start string)
	visit = func(start string) {
		state[start] = TemporaryMark
		for _, item := range m[start] {
			if state[item] == Unvisited {
				visit(item)
			}
		}
		order = append(order, start)
	}
	visitAll = func() {
		for item := range prereqs {
			if state[item] == Unvisited {
				visit(item)
			}
		}
	}
	visitAll()
	return order
}

package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)

	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	outline(nil, doc)
	frequency := make(map[string]int)
	tagFrequency(frequency, doc)

	fmt.Println("\n\nHTML Tag frequency")
	for name, count := range frequency {
		fmt.Printf("%s\t\t%d\n", name, count)
	}
}

func tagFrequency(frequencyMap map[string]int, n *html.Node) {
	if n.Type == html.ElementNode {
		frequencyMap[n.Data] +=1
	}
	for i := n.FirstChild; i != nil; i = i.NextSibling {
		tagFrequency(frequencyMap, i)
	}
}

func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data) // p ush tag
		fmt.Println(stack)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}

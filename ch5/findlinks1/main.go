package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

var srcKey = map[string]string {"a": "href", "img": "src", "link": "href", "script": "src"}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && srcKey[n.Data] != "" {
		for _, a := range n.Attr {
			if a.Key == srcKey[n.Data] {
				links = append(links, a.Val)

			}

		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

func visitRecursively(links []string, n *html.Node) []string {
	// change the findlinks program to traverse the n.FirstChild linked list using recursive calls to visit instead of a loop
	if n == nil {
		return links
	}
	if n.Type == html.ElementNode && srcKey[n.Data] != "" {
		for _, a := range n.Attr {
			if a.Key == srcKey[n.Data] {
				links = append(links, a.Val)
			}
		}
	}
	links = visitRecursively(links, n.NextSibling)
	links = visitRecursively(links, n.FirstChild)
	return links
}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlink1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visitRecursively(nil, doc) {
		fmt.Println(link)
	}
}

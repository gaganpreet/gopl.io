package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
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
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
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

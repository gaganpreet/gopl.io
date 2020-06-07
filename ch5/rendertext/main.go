package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

var blacklistedHTMLTags []string = []string{"script", "style"}

func printText(n *html.Node, prefix string) {
	if n.Type == html.ElementNode {
		for _, tag := range blacklistedHTMLTags {
			if n.Data == tag {
				return
			}
		}
	}
	if n.Type == html.TextNode {
		fmt.Println(prefix, n.Data)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		printText(c, fmt.Sprintf(" %s", prefix))
	}
}

func main() {
	html, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while parsing html: %s", err)
	}
	printText(html, "")
}

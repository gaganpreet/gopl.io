package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	if n.Type == html.ElementNode {
		if n.Data == "script" || n.Data == "style" {
			return
		}
	}
	if n.Type == html.TextNode {
		words += len(strings.Split(n.Data, " "))
	} else if n.Type == html.ElementNode && n.Data == "img" {
		images++
	}
	for i := n.FirstChild; i != nil; i = i.NextSibling {
		childWordCount, childImageCount := countWordsAndImages(i)
		words += childWordCount
		images += childImageCount
	}
	return
}

func main() {
	for _, url := range os.Args[1:] {
		words, images, _ := CountWordsAndImages(url)
		fmt.Println(words, images)
	}
}

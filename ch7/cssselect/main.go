package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

type CSSSelector struct {
	elementType       string
	classSelector     string
	idSelector        string
	attributeSelector string
	attributeValue    string
}

func main() {
	dec := xml.NewDecoder(os.Stdin)
	var stack []string // stack of element names
	matches := false
	selector := CSSSelector{elementType: "a"}
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack, tok.Name.Local) // push
			if isElementMatching(selector, tok) {
				matches = true
			}
		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop
			matches = false
		case xml.CharData:
			if matches {
				fmt.Printf("%s: %q\n", strings.Join(stack, " "), tok)
			}
		}
	}
}

func isElementMatching(selector CSSSelector, node xml.StartElement) bool {
	if selector.elementType != "" && selector.elementType != node.Name.Local {
		return false
	}
	if selector.idSelector == "" && selector.classSelector == "" && selector.attributeSelector == "" {
		return true
	}
	for _, attr := range node.Attr {
		if attr.Name.Local == "id" && selector.idSelector != "" && attr.Value == selector.idSelector {
			return true
		} else if attr.Name.Local == "class" && selector.classSelector != "" && attr.Value == selector.classSelector {
			return true
		} else if selector.attributeSelector != "" {
			if attr.Name.Local == selector.attributeSelector {
				return true
			}
			if selector.attributeValue == "" && attr.Value == selector.attributeValue {
				return true
			}
		}
	}
	return false
}

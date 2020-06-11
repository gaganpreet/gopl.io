package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopl.io/ch5/links"
)

func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func crawl(url string)[] string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func crawlFs(dir string)[] string {
	fmt.Println(dir)
	dirEntries, err := ioutil.ReadDir(dir) 
	if err != nil {
		return nil
	}

	var childDirs []string
	for _, dirEntry := range dirEntries {
		if dirEntry.IsDir() {
			childDirs = append(childDirs, fmt.Sprintf("%s/%s", dir, dirEntry.Name()))
		}
	}
	return childDirs
}

func main() {
	breadthFirst(crawlFs, os.Args[1:])
}

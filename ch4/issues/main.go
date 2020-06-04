package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gopl.io/ch3/ch4/github"
)

func timeSince(date time.Time) string {
	MONTH, _ := time.ParseDuration(fmt.Sprintf("%dh", 24*30))
	YEAR, _ := time.ParseDuration(fmt.Sprintf("%dh", 24*365))

	since := time.Since(date)

	if since <= MONTH {
		return "less than a month"
	} else if since <= YEAR {
		return "less than a year"
	}
	return "more than a year"
}

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d issues:\n", result.TotalCount)
	issueAgeCategories := make(map[string]int)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
		issueAgeCategories[timeSince(item.CreatedAt)]++
	}

	fmt.Printf("\nAge categories\n")
	for ageCategory, count := range issueAgeCategories {
		fmt.Printf("%s\t%d\n", ageCategory, count)
	}
}

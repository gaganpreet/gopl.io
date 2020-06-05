package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

const cacheDir string = "./xkcdcache"

type ComicInfo struct {
	Month      string
	Num        int
	Link       string
	Year       string
	News       string
	SafeTitle  string `json:"safe_title"`
	Transcript string
	Alt        string
	Img        string
	Title      string
	Day        string
}

func getLatestComic() (*ComicInfo, error) {
	latestComicURL := "http://xkcd.com/info.0.json"
	comic, err := fetchComicFromURL(latestComicURL)
	return comic, err
}

func getComic(id int) (*ComicInfo, error) {
	comicURL := fmt.Sprintf("http://xkcd.com/%d/info.0.json", id)
	return fetchComicFromURL(comicURL)
}

func fetchComicFromURL(url string) (*ComicInfo, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		resp.Body.Close()
		return nil, fmt.Errorf("failed to fetch %s with status code %d", url, resp.StatusCode)
	}

	var result ComicInfo
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}

func fetchAndCacheComic(id int, query string) {
	cacheLocation := fmt.Sprintf("%s/%d.json", cacheDir, id)
	file, err := os.Open(cacheLocation)
	var comic *ComicInfo = new(ComicInfo)
	if err != nil {
		comic, err = getComic(id)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error while fetching %d: %s\n", id, err)
			return
		}
		writeFile, err := os.Create(cacheLocation)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error while writing to cache <%s> for %d: %s\n", cacheLocation, id, err)
			return
		}
		json.NewEncoder(writeFile).Encode(comic)
		writeFile.Close()
		return
	} else {
		json.NewDecoder(file).Decode(comic)
	}
	file.Close()
	if comic != nil && (strings.Contains(strings.ToLower(comic.Transcript), query) ||
		strings.Contains(strings.ToLower(comic.Alt), query)) {
		fmt.Printf("[#%d]\thttps://xkcd.com/%d\t%s\n", comic.Num, comic.Num, comic.Alt)
	}
}

func cacheAndSearchComics(query string) error {
	comic, err := getLatestComic()
	if err != nil {
		return err
	}
	latest := comic.Num
	for i := 1; i <= latest; i++ {
		fetchAndCacheComic(i, query)
	}
	return nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Need a search argument")
	}
	query := os.Args[1]
	err := cacheAndSearchComics(query)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed with %s", err)
	}
}

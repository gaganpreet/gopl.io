package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
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

func fetchAndCacheComic(id int, ch chan<- string) {
	cacheLocation := fmt.Sprintf("%s/%d.json", cacheDir, id)
	file, err := os.Open(cacheLocation)
	if err != nil {
		comic, err := getComic(id)
		if err != nil {
			ch <- fmt.Sprintf("error while fetching %d: %s", id, err)
			return
		}
		writeFile, err := os.Create(cacheLocation)
		if err != nil {
			ch <- fmt.Sprintf("error while writing to cache <%s> for %d: %s", cacheLocation, id, err)
			return
		}
		json.NewEncoder(writeFile).Encode(comic)
		writeFile.Close()
		ch <- fmt.Sprintf("Cached comic %d in %s", id, cacheLocation)
		return
	}
	file.Close()
	ch <- fmt.Sprintf("Found existing cache for %d in %s", id, cacheLocation)
}

func cacheComics() error {
	comic, err := getLatestComic()
	if err != nil {
		return err
	}
	latest := comic.Num
	ch := make(chan string, 10)
	for i := 1; i <= latest; i++ {
		go fetchAndCacheComic(i, ch)
	}
	for i := 1; i <= latest; i++ {
		fmt.Println(<-ch)
	}
	return nil
}

func main() {
	if os.Args[1] == "init" {
		err := cacheComics()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed with %s", err)
		}
	}
}

// Exercies 7.9, print tracks in html
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"sort"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

type byArtist []*Track

func (x byArtist) Len() int           { return len(x) }
func (x byArtist) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }
func (x byArtist) Less(i, j int) bool { return x[i].Artist < x[j].Artist }

type customSort struct {
	t    []*Track
	less func(x, y *Track) bool
}

func (x customSort) Len() int { return len(x.t) }
func (x customSort) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }
func (x customSort) Swap(i, j int) { x.t[i], x.t[j] = x.t[j], x.t[i] }

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

var trackList = template.Must(template.New("tracklist").Parse(`
<table>
<tr style='text-align: left'>
<th><a href="?sort=title">Title</th>
<th><a href="?sort=artist">Artist</th>
<th><a href="?sort=year">Year</th>
<th><a href="?sort=album">Album</th>
<th><a href="?sort=length">Length</th>
{{ range $ }}
<tr>
<td>{{ .Title }}</td>
<td>{{ .Artist }}</td>
<td>{{ .Year }}</td>
<td>{{ .Album }}</td>
<td>{{ .Length }}</td>
{{ end }}
</table>
`))

func tracksHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		return
	}
	fmt.Fprintf(w, "<html>")
	sortKey := r.FormValue("sort")
	sort.Sort(customSort{tracks, func(x, y *Track) bool {
		if sortKey == "title" {
			return x.Title < y.Title
		}
		if sortKey == "artist" {
			return x.Artist < y.Artist
		}
		if sortKey == "album" {
			return x.Album < y.Album
		}
		if sortKey == "year" {
			return x.Year < y.Year
		}
		if sortKey == "length" {
			return x.Length < y.Length
		}
		return false
	}})
	if err := trackList.Execute(w, tracks); err != nil {
		fmt.Fprintf(w, "%s", err)
	}
	fmt.Fprintf(w, "</html>")
}

func main() {
	http.HandleFunc("/", tracksHandler)
	log.Fatal(http.ListenAndServe("localhost:8005", nil))
}

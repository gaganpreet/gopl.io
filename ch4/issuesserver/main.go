package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"gopl.io/ch4/github"
)

var issueList = template.Must(template.New("issuelist").Parse(`
<h1>{{.TotalCount}} issues </h1>
<table>
<tr style='text-align: left'>
 <th>#</th>
 <th>State</th>
 <th>User</th>
 <th>Title</th>
</tr>
{{ range .Items  }}
<tr>
<td><a href='{{ .HTMLURL }}'>{{.Number}}</a></td>
<Td>{{.State}}</td>
<td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
<td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
`))

var inputBox = template.Must(template.New("inputbox").Parse(`
<form action="/">
<input type="text" name="query" value="{{ . }}">
<input type="submit" value="Submit">
</form>
`))

func issueListHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		return
	}
	fmt.Fprintf(w, "<html>")
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	query := r.FormValue("query")
	if len(query) == 0 {
		fmt.Fprintf(w, "Enter a search term")
	}
	if err := inputBox.Execute(w, query); err != nil {
		fmt.Fprintf(w, "%s", err)
	}

	if len(query) > 0 {
		result, err := github.SearchIssues([]string{query})
		if err != nil {
			fmt.Fprintf(w, "%s", err)
		}
		if err := issueList.Execute(w, result); err != nil {
			fmt.Fprintf(w, "%s", err)
		}
	}
	fmt.Fprintf(w, "</html>")
}

func main() {
	http.HandleFunc("/", issueListHandler)
	log.Fatal(http.ListenAndServe("localhost:8005", nil))
}

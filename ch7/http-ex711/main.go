package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var mu sync.Mutex

var listTemplate = template.Must(template.New("listProducts").Parse(`
<html>
<head>
<title>Shopping</title>
</head>
<body>
<table>
<tr style='text-align: left'>
	<th>Item</th>
	<th>Price</th>
</tr>
{{ range $key, $value := . }}
<tr>
	<td>{{ $key }}</td>
	<td>{{ $value }}</td>
</tr>
{{ end }}
</table>
</body>
`))

func main() {
	db := database{"shoes": 50, "socks": 5}
	// Not REST, but whatever
	http.HandleFunc("/add", db.addOrUpdateItem)    // CREATE
	http.HandleFunc("/update", db.addOrUpdateItem) // UPDATE
	http.HandleFunc("/delete", db.deleteItem)      // DELETE
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type database map[string]dollars

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

func (db database) list(w http.ResponseWriter, req *http.Request) {
	if err := listTemplate.Execute(w, db); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

func (db database) addOrUpdateItem(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	item := query.Get("item")
	price := query.Get("price")
	priceInt, err := strconv.Atoi(price)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "invalid price for %q: %q\n", item, price)
		return
	}
	mu.Lock()
	db[item] = dollars(priceInt)
	mu.Unlock()
	fmt.Fprintf(w, "successfully added %q", item)
}

func (db database) deleteItem(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	mu.Lock()
	itemName := query.Get("item")
	if _, ok := db[itemName]; !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "item %q not found\n", itemName)
		return
	}
	delete(db, itemName)
	mu.Unlock()
	fmt.Fprintf(w, "successfully removed %q", itemName)
}

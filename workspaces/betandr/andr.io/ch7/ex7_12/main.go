// Change the handler for `/list` to print its output as an HTML table, not text.
// You may find the `html/template` package (§4.6) useful.
package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"
)

// curl -X GET http://localhost:8000/list
// curl -X GET http://localhost:8000/price\?item\=socks
// curl -X GET http://localhost:8000/update?item=socks&price=6
// curl -X GET http://localhost:8000/price\?item\=socks
// curl -X GET http://localhost:8000/create\?item\=sandals\&price\=20
// curl -X GET http://localhost:8000/show\?item\=sandals
// curl -X GET http://localhost:8000/delete\?item\=sandals

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/show", db.read)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/create", db.create)
	http.HandleFunc("/delete", db.delete)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type pounds float32

func (p pounds) String() string { return fmt.Sprintf("£%.2f", p) }

type database map[string]pounds

func (db database) list(w http.ResponseWriter, req *http.Request) {
	render(w, db)
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if price, ok := db[item]; ok {
		fmt.Fprintf(w, "%s\n", price)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

func (db database) create(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if _, ok := db[item]; ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "item exists: %s\n", item)
		return
	}
	p := req.URL.Query().Get("price")
	price, err := strconv.ParseFloat(p, 10)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "invalid price: %s\n", p)
		return
	}
	db[item] = pounds(price)
	render(w, db)
}

func (db database) read(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")

	if _, ok := db[item]; ok {
		fmt.Fprintf(w, "%s: %s\n", item, db[item])
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	p := req.URL.Query().Get("price")
	price, err := strconv.ParseFloat(p, 10)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "invalid price: %s\n", p)
		return
	}

	if _, ok := db[item]; ok {
		db[item] = pounds(price)
		render(w, db)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

func (db database) delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")

	if _, ok := db[item]; ok {
		delete(db, item)
		if _, ok := db[item]; !ok {
			render(w, db)
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

func render(out io.Writer, db map[string]pounds) {
	templ := `
	<html>
	<style>
		table, th, td {
			width: 200;
			border: 1px solid black;
		}
		td {
			padding: 10px;
		}
	</style>
	<table>
		<th>Item</th>
		<th>Price</th>
		{{ range $key, $value := . }}
		<tr>
		<td>{{ $key }}</td>
		<td>{{ $value }}</td>
		</tr>
		{{end}}
	</table>
	</html>
	`
	tracksList := template.Must(template.New("items").Parse(templ))

	if err := tracksList.Execute(out, db); err != nil {
		log.Fatal(err)
	}
}

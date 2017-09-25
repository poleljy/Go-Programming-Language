package ch7

import (
	"fmt"
	"log"
	"net/http"
)

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s:%s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item:%q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

func TestHTTP2() {
	db := database{"shoes": 50, "socks": 5}

	//mux := http.NewServeMux()
	//mux.Handle("/list", http.HandlerFunc(db.list))
	//mux.Handle("/price", http.HandlerFunc(db.price))

	http.HandleFunc("/list", http.HandlerFunc(db.list))
	http.HandleFunc("/price", http.HandlerFunc(db.price))

	log.Fatal(http.ListenAndServe("localhost:8080", db))
}

package main

import (
	"fmt"
	"log"
	"net/http"
)

func list(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "list")
}

func detail(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "detail")
}

func main() {

	//mux := http.NewServeMux()
	//mux.Handle("/list", http.HandlerFunc(list))
	//mux.Handle("/detail", http.HandlerFunc(detail))
	//mux.HandleFunc("/list", list)
	//mux.HandleFunc("/detail", detail)

	//log.Fatal(http.ListenAndServe("localhost:8080", mux))

	http.HandleFunc("/list", list)
	http.HandleFunc("/detail", detail)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

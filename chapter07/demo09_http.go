package main

import (
	"fmt"
	"log"
	"net/http"
)

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for k, v := range db {
		fmt.Fprintf(w, "%s: %s\n", k, v)
	}
}

func main() {
	db := database{
		"shone": 50,
		"socks": 5,
	}
	log.Fatal(http.ListenAndServe("localhost:8080", db))
}

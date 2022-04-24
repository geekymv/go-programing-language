package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.RWMutex
var count int

func main() {

	http.HandleFunc("/hello", handlerRequest)
	http.HandleFunc("/counter", counter)
	log.Fatalln(http.ListenAndServe("localhost:8000", nil))

}

func handlerRequest(w http.ResponseWriter, r *http.Request) {
	log.Println("handler request")
	mu.Lock()
	count++
	mu.Unlock()

	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request) {
	log.Println("counter")
	mu.RLocker().Lock()
	fmt.Fprintf(w, "Counter:%d", count)
	mu.RLocker().Unlock()

}

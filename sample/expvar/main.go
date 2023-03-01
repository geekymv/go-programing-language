package main

import (
	"expvar"
	"log"
	"net"
	"net/http"
	"time"
)

// Go标准库 expvar

func main() {

	mux := http.NewServeMux()

	mux.Handle("/debug/vars", expvar.Handler())
	start := time.Now()
	expvar.Publish("Uptime", expvar.Func(func() interface{} {
		return time.Since(start).Seconds()
	}))

	mux.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		resp.Header().Set("Content-Type", "application/json; charset=utf-8")
		resp.WriteHeader(http.StatusNotFound)
		resp.Write([]byte("not found"))
	})

	server := &http.Server{
		Handler: mux,
	}
	addr := ":8899"
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln(err)
	}

	done := make(chan bool)

	go func() {
		log.Printf("http listen:%v\n", addr)
		err = server.Serve(listen)
		if err != nil {
			log.Fatalln(err)
		}
	}()

	<-done

}

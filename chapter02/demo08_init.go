package main

import (
	"fmt"
	"log"
	"os"
)

var cwd string

func init() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("os.Getwd failed:%v", err)
	}
	log.Printf("Working directory = %s", cwd)
}

func main() {
	fmt.Println("main:", cwd)
}

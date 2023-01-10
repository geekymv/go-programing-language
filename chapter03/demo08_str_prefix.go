package main

import (
	"fmt"
	"strings"
)

func main() {
	oid := "https://www.test.com"
	if strings.HasPrefix(oid, "http") || strings.HasPrefix(oid, "https") {
		fmt.Println("true")
	}
}

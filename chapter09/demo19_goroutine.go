package main

import (
	"fmt"
)

func main() {
	//runtime.GOMAXPROCS(1)
	for {
		go fmt.Print(0)
		fmt.Print(1)
	}

}

package main

import (
	"fmt"
	"os"
)

func main() {

	f, err := os.Open("/a/b/c")
	fmt.Printf("err:%T\n", err)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("file not exist")
		}
		return
	}

	fmt.Println("file name =", f.Name())

}

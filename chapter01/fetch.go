package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {

	for _, url := range os.Args[1:] {

		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("fetch %s err:%v\n", url, err)
			os.Exit(1)
		}

		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()

		if err != nil {
			fmt.Printf("fetch %s read err:%v", url, err)
			os.Exit(1)
		}

		fmt.Println(string(b))
	}

}

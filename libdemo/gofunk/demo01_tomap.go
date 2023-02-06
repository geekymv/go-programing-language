package main

import (
	"fmt"

	"github.com/thoas/go-funk"
)

type Foo struct {
	ID        int
	FirstName string `tag_name:"tag 1"`
	LastName  string `tag_name:"tag 2"`
	Age       int    `tag_name:"tag 3"`
}

func main() {
	f := &Foo{
		ID:        1,
		FirstName: "Gilles",
		LastName:  "Fabio",
		Age:       70,
	}

	b := &Foo{
		ID:        1,
		FirstName: "Florent",
		LastName:  "Messa",
		Age:       80,
	}

	results := []*Foo{f, b}
	// slice/array 转成 map
	mapping, ok := funk.ToMap(results, "ID").(map[int]*Foo)
	if !ok {
		fmt.Println("not ok")
	} else {
		for k, v := range mapping {
			fmt.Println(k, v)
		}
	}
}

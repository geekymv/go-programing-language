package main

import "fmt"

type Values map[string][]string

func (v Values) Get(key string) string {
	if vs := v[key]; len(vs) > 0 {
		return vs[0]
	}
	return ""
}

func (v Values) Add(key, value string) {
	v[key] = append(v[key], value)
}

func main() {
	m := Values{}
	m.Add("a", "b")
	m.Add("a", "c")
	fmt.Println(m)

	m = nil
	m.Add("c", "d") // panic
}

package main

import "fmt"

func nonempty(s []string) []string {
	// 0个元素，和s共用底层数组
	out := s[:0]

	for _, v := range s {
		if v != "" {
			out = append(out, v)
		}
	}

	return out
}

func main() {

	s := []string{"a", "", "c"}
	s = nonempty(s)
	fmt.Println(s)
}

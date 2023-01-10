package main

import "fmt"

func appendSliceWithinCap(s []string) {
	s = append(s, "two")
	s[0] = "appendSliceWithinCap"
}
func appendSliceOverCap(s []string) {
	s = append(s, "two")
	s = append(s, "three")
	s[0] = "appendSliceOverCap"
}
func main() {
	fmt.Println("hello main")
	s := make([]string, 1, 2)
	s[0] = "one"
	fmt.Println(s) // ["one"]
	appendSliceWithinCap(s)
	fmt.Println(s) // ["appendSliceWithinCap"]
	appendSliceOverCap(s)
	fmt.Println(s) // ["appendSliceWithinCap"]
}

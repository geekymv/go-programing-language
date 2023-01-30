package main

import "fmt"

func main() {

	m := make(map[string]string)
	foo(m)
	fmt.Println(m, len(m))

	tidList := make([]int, 2, 2)
	bar(tidList)
	fmt.Println(tidList)
}

func bar(tidList []int) {
	//tidList[0] = "0"
	//tidList[1] = "1"
}

func foo(m map[string]string) {
	m["0"] = "0"
	m["1"] = "0"
	m["2"] = "0"
	m["3"] = "0"
	m["4"] = "0"
	m["5"] = "0"
	m["6"] = "0"
	m["7"] = "0"
	m["8"] = "0"
	m["9"] = "0"
	m["10"] = "0"
}

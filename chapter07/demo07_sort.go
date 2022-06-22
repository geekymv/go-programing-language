package main

import (
	"fmt"
	"sort"
)

// sort.Interface 定义了排序的规则

type StringSlice []string

func (s StringSlice) Len() int {
	return len(s)
}

func (s StringSlice) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s StringSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {

	names := []string{
		"tom",
		"jack",
		"tony",
	}

	sort.Sort(StringSlice(names))
	fmt.Println(names)
	//sort.Strings(names)

	// 组合模式，对自定义的排序规则进行修改，反回一个新的规则（逆序）
	sort.Sort(sort.Reverse(StringSlice(names)))

	fmt.Println(names)
}

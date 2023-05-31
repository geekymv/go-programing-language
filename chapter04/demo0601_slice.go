package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {

	languages := []string{"Go", "Java", "C"}

	languages, err := RemoveByIndex(languages, 0)
	if err != nil {
		log.Fatalf("remove by index err:%v", err)
		return
	}
	fmt.Println(languages)
}

// RemoveByIndex 根据索引删除元素
func RemoveByIndex(strs []string, index int) ([]string, error) {
	if index >= len(strs) || index < 0 {
		return nil, errors.New("out of index")
	}
	strs = append(strs[0:index], strs[index+1:]...)
	return strs, nil
}

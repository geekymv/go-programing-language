package main

import (
	"errors"
	"fmt"
)

// op 多返回值
func op(x, y int) (int, error) {
	if y == 0 {
		return -1, errors.New("除数不能为0")
	}
	return x / y, nil
}

func main() {

	if r, err := op(10, 0); err != nil {
		fmt.Println("err=", err.Error())
	} else {
		fmt.Println("r=", r)
	}
}

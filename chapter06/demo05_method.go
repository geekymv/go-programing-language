package main

import (
	"fmt"
	"time"
)

type Rocket struct {
}

func (r *Rocket) Launch() {
	fmt.Println("launch...")
}

func main() {

	r := new(Rocket)

	// 方法变量
	time.AfterFunc(time.Second*5, r.Launch)

	time.Sleep(15 * time.Second)

}

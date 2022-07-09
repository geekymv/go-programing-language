package main

import (
	"fmt"
	"time"
)

func launch() {
	fmt.Println("launch")
}

func main() {
	tick := time.Tick(1 * time.Second)
	for count := 10; count > 0; count-- {
		fmt.Println(count)
		<-tick
	}

	launch()
}

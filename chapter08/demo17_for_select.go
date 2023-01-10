package main

import (
	"log"
	"time"
)

func main() {

	workTicker := time.NewTicker(time.Second * 10)
	refreshTicker := time.NewTicker(time.Second * 2)
	i := 0

	for {
		select {
		case <-workTicker.C:
			log.Println("work")
			i++
		case <-refreshTicker.C:
			log.Println("refresh")
		}

		// select 执行了其中一个 case 之后下面代码才会执行
		log.Println("i=", i)
	}
}

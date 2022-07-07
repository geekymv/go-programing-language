package main

import "fmt"

func main() {

	worklist := make(chan []string)
	go func() {
		worklist <- []string{"a", "b"}
		close(worklist)
	}()

	// for range 遍历 channel 需要先将 channel close，否则会发生 deadlock。
	for list := range worklist {
		fmt.Println(list)
	}

}

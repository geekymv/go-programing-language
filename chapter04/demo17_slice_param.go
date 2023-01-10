package main

import "fmt"

type SliceParam struct {
	ID int
}

func main() {

	sp := make([]*SliceParam, 0, 2)
	for i := 0; i < 2; i++ {
		sp = append(sp, &SliceParam{
			ID: i,
		})
	}

	appendSlice(sp)

	fmt.Println("---------------------")
	fmt.Println(len(sp), cap(sp))
	for _, v := range sp {
		fmt.Println(v.ID)
	}

}

func appendSlice(sp []*SliceParam) {
	for _, v := range sp {
		v.ID = v.ID + 1
	}

	sp = append(sp, &SliceParam{
		ID: 3,
	})

	fmt.Println(len(sp), cap(sp))
	for _, v := range sp {
		fmt.Println(v.ID)
	}
}
